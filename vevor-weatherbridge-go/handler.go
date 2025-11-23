package main

import (
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"time"
)

// WeatherHandler handles incoming weather station data.
type WeatherHandler struct {
	cfg  *Config
	mqtt *MQTTClient
	wu   *WUForwarder
}

// NewWeatherHandler creates a new weather handler.
func NewWeatherHandler(cfg *Config, mqtt *MQTTClient, wu *WUForwarder) *WeatherHandler {
	return &WeatherHandler{
		cfg:  cfg,
		mqtt: mqtt,
		wu:   wu,
	}
}

// ServeHTTP handles the weather station update endpoint.
func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Received weather update request", "path", r.URL.Path, "query", r.URL.RawQuery)

	// Parse query parameters
	query := r.URL.Query()

	// Parse timestamp
	var measuredTime string
	if dateutc := query.Get("dateutc"); dateutc != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateutc)
		if err == nil {
			localTime := parsedTime.In(h.cfg.Timezone)
			measuredTime = localTime.Format(time.RFC3339)
		} else {
			slog.Warn("Failed to parse dateutc", "value", dateutc, "error", err)
			measuredTime = time.Now().In(h.cfg.Timezone).Format(time.RFC3339)
		}
	} else {
		measuredTime = time.Now().In(h.cfg.Timezone).Format(time.RFC3339)
	}

	// Process each sensor
	publishedCount := 0
	for _, sensor := range SensorDefinitions {
		rawValue := query.Get(sensor.QueryParam)
		if rawValue == "" {
			continue
		}

		// Parse the value
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			slog.Warn("Failed to parse sensor value", "sensor", sensor.ID, "value", rawValue, "error", err)
			continue
		}

		// Convert value based on sensor type and units
		convertedValue := h.convertValue(&sensor, value)

		// Format the value for publishing
		stateValue := h.formatValue(&sensor, convertedValue)

		// Publish sensor config
		if err := h.mqtt.PublishSensorConfig(&sensor); err != nil {
			slog.Error("Failed to publish sensor config", "sensor", sensor.ID, "error", err)
			continue
		}

		// Publish sensor state
		if err := h.mqtt.PublishSensorState(sensor.ID, stateValue); err != nil {
			slog.Error("Failed to publish sensor state", "sensor", sensor.ID, "error", err)
			continue
		}

		// Build and publish attributes
		attrs := map[string]interface{}{
			"measured_on": measuredTime,
		}

		// Add cardinal direction for wind direction sensor
		if sensor.ID == "wind_direction" {
			attrs["cardinal"] = DegreesToCardinal(value)
		}

		if err := h.mqtt.PublishSensorAttributes(sensor.ID, attrs); err != nil {
			slog.Error("Failed to publish sensor attributes", "sensor", sensor.ID, "error", err)
			continue
		}

		publishedCount++
		slog.Debug("Published sensor data", "sensor", sensor.ID, "value", stateValue)
	}

	slog.Info("Processed weather update", "sensors_published", publishedCount)

	// Forward to Weather Underground if enabled
	if h.cfg.WUForward && h.wu != nil {
		go h.wu.Forward(r.URL.Query())
	}

	// Always return success to the weather station
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "success")
}

// convertValue applies unit conversion based on sensor type and configured units.
func (h *WeatherHandler) convertValue(sensor *SensorDefinition, value float64) float64 {
	if h.cfg.IsMetric() {
		switch sensor.QueryParam {
		case "tempf", "dewptf":
			return FToC(value)
		case "baromin":
			return InHgToHPa(value)
		case "windspeedmph", "windgustmph":
			return MphToKmh(value)
		case "rainin", "dailyrainin":
			return InchToMm(value)
		}
	} else {
		// Imperial: just round appropriately
		switch sensor.QueryParam {
		case "tempf", "dewptf", "baromin", "windspeedmph", "windgustmph":
			return roundTo(value, 1)
		case "rainin", "dailyrainin":
			return roundTo(value, 2)
		}
	}
	return value
}

// formatValue formats the value as a string for MQTT publishing.
func (h *WeatherHandler) formatValue(sensor *SensorDefinition, value float64) string {
	// Handle special cases
	switch sensor.QueryParam {
	case "humidity", "UV", "winddir":
		// These are integers or have no decimals
		if math.IsNaN(value) {
			return ""
		}
		return strconv.FormatFloat(value, 'f', -1, 64)
	case "rainin", "dailyrainin":
		if !h.cfg.IsMetric() {
			// Imperial rain uses 2 decimal places
			return strconv.FormatFloat(value, 'f', 2, 64)
		}
		return strconv.FormatFloat(value, 'f', 1, 64)
	default:
		// Default to 1 decimal place
		return strconv.FormatFloat(value, 'f', 1, 64)
	}
}
