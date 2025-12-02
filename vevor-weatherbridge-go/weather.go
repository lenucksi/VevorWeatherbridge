// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (C) 2025 Lenucksi
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	// WUHost is the Weather Underground host for updates.
	WUHost = "rtupdate.wunderground.com"
	// WUPath is the endpoint path for weather updates.
	WUPath = "/weatherstation/updateweatherstation.php"
)

// WUForwarder forwards weather data to Weather Underground.
type WUForwarder struct {
	cfg      *Config
	client   *http.Client
	resolver *net.Resolver
}

// NewWUForwarder creates a new Weather Underground forwarder.
func NewWUForwarder(cfg *Config) *WUForwarder {
	// Create a custom resolver using Google DNS to bypass local DNS redirect
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: 5 * time.Second}
			// Use Google DNS servers
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				Resolver:  resolver,
			}).DialContext,
		},
	}

	return &WUForwarder{
		cfg:      cfg,
		client:   client,
		resolver: resolver,
	}
}

// Forward sends the weather data to Weather Underground.
func (w *WUForwarder) Forward(params url.Values) {
	// Resolve the real IP of WU to bypass local DNS redirect
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ips, err := w.resolver.LookupIP(ctx, "ip4", WUHost)
	if err != nil || len(ips) == 0 {
		slog.Error("Failed to resolve Weather Underground IP", "error", err)
		return
	}

	wuIP := ips[0].String()
	slog.Debug("Resolved Weather Underground IP", "ip", wuIP)

	// Clone params and override credentials if configured
	forwardParams := url.Values{}
	for k, v := range params {
		forwardParams[k] = v
	}

	if w.cfg.WUUsername != "" {
		forwardParams.Set("ID", w.cfg.WUUsername)
	}
	if w.cfg.WUPassword != "" {
		forwardParams.Set("PASSWORD", w.cfg.WUPassword)
	}

	// Build the request URL using the resolved IP
	wuURL := fmt.Sprintf("http://%s%s?%s", wuIP, WUPath, forwardParams.Encode())

	// Create request with Host header for virtual hosting
	req, err := http.NewRequest("GET", wuURL, nil)
	if err != nil {
		slog.Error("Failed to create WU request", "error", err)
		return
	}
	req.Host = WUHost

	// Send request
	resp, err := w.client.Do(req)
	if err != nil {
		slog.Error("Failed to forward to Weather Underground", "error", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusOK {
		slog.Info("Successfully forwarded to Weather Underground")
	} else {
		slog.Warn("Weather Underground returned non-OK status", "status", resp.StatusCode)
	}
}
