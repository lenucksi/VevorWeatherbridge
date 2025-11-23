package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load configuration
	cfg := LoadConfig()

	// Setup structured logging
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	})
	slog.SetDefault(slog.New(logHandler))

	slog.Info("Starting VEVOR Weather Station Bridge (Go)",
		"version", Version,
		"device_name", cfg.DeviceName,
		"device_id", cfg.DeviceID,
		"units", cfg.Units,
		"timezone", cfg.Timezone.String(),
	)

	// Connect to MQTT broker
	slog.Info("Connecting to MQTT broker", "host", cfg.MQTTHost, "port", cfg.MQTTPort)
	mqttClient, err := NewMQTTClient(cfg)
	if err != nil {
		slog.Error("Failed to connect to MQTT broker", "error", err)
		os.Exit(1)
	}
	defer mqttClient.Close()

	// Create Weather Underground forwarder if enabled
	var wuForwarder *WUForwarder
	if cfg.WUForward {
		slog.Info("Weather Underground forwarding enabled")
		wuForwarder = NewWUForwarder(cfg)
	}

	// Create HTTP handler
	handler := NewWeatherHandler(cfg, mqttClient, wuForwarder)

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.Handle("/weatherstation/updateweatherstation.php", handler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		if mqttClient.IsConnected() {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("MQTT disconnected"))
		}
	})

	server := &http.Server{
		Addr:         ":80",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		slog.Info("HTTP server starting", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("HTTP server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan

	slog.Info("Received shutdown signal", "signal", sig)

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error during server shutdown", "error", err)
	}

	slog.Info("Server stopped")
}
