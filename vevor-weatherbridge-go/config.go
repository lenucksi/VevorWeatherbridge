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
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all application configuration from environment variables.
type Config struct {
	// Logging
	LogLevel slog.Level

	// MQTT configuration
	MQTTHost     string
	MQTTPort     int
	MQTTUser     string
	MQTTPassword string
	MQTTPrefix   string

	// Device identification
	DeviceID           string
	DeviceName         string
	DeviceManufacturer string
	DeviceModel        string

	// Timezone
	Timezone *time.Location

	// Units (metric or imperial)
	Units string

	// Weather Underground forwarding
	WUForward  bool
	WUUsername string
	WUPassword string
}

// LoadConfig loads configuration from environment variables with defaults.
func LoadConfig() *Config {
	cfg := &Config{
		LogLevel:           parseLogLevel(getEnv("LOG_LEVEL", "INFO")),
		MQTTHost:           getEnv("MQTT_HOST", "localhost"),
		MQTTPort:           getEnvInt("MQTT_PORT", 1883),
		MQTTUser:           getEnv("MQTT_USER", ""),
		MQTTPassword:       getEnv("MQTT_PASSWORD", ""),
		MQTTPrefix:         getEnv("MQTT_PREFIX", "homeassistant"),
		DeviceName:         getEnv("DEVICE_NAME", "Weather Station"),
		DeviceManufacturer: getEnv("DEVICE_MANUFACTURER", "VEVOR"),
		DeviceModel:        getEnv("DEVICE_MODEL", "7-in-1 Weather Station"),
		Units:              strings.ToLower(getEnv("UNITS", "metric")),
		WUForward:          getEnvBool("WU_FORWARD", false),
		WUUsername:         getEnv("WU_USERNAME", ""),
		WUPassword:         getEnv("WU_PASSWORD", ""),
	}

	// Derive DeviceID from DeviceName (lowercase, spaces to underscores)
	deviceID := getEnv("DEVICE_ID", "")
	if deviceID == "" {
		deviceID = strings.ToLower(cfg.DeviceName)
		deviceID = strings.ReplaceAll(deviceID, " ", "_")
	}
	cfg.DeviceID = deviceID

	// Parse timezone
	tzName := getEnv("TZ", "Europe/Berlin")
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		slog.Warn("Invalid timezone, using UTC", "timezone", tzName, "error", err)
		loc = time.UTC
	}
	cfg.Timezone = loc

	// Validate units
	if cfg.Units != "metric" && cfg.Units != "imperial" {
		slog.Warn("Invalid units, defaulting to metric", "units", cfg.Units)
		cfg.Units = "metric"
	}

	return cfg
}

// getEnv returns environment variable value or default.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt returns environment variable as int or default.
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

// getEnvBool returns environment variable as bool or default.
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		value = strings.ToLower(value)
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}

// parseLogLevel converts string log level to slog.Level.
func parseLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARNING", "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// IsMetric returns true if metric units are configured.
func (c *Config) IsMetric() bool {
	return c.Units == "metric"
}
