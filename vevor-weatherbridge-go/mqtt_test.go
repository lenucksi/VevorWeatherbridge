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
	"encoding/json"
	"testing"
	"time"
)

func TestDiscoveryPayloadJSON(t *testing.T) {
	cfg := &Config{
		MQTTPrefix:         "homeassistant",
		DeviceID:           "weather_station",
		DeviceName:         "Weather Station",
		DeviceManufacturer: "VEVOR",
		DeviceModel:        "7-in-1 Weather Station",
		Units:              "metric",
	}

	payload := DiscoveryPayload{
		Name:                "Weather Station Temperature",
		StateTopic:          "homeassistant/sensor/weather_station_temperature/state",
		UniqueID:            "weather_station_temperature",
		DeviceClass:         "temperature",
		UnitOfMeasurement:   "°C",
		AvailabilityTopic:   "homeassistant/sensor/weather_station/availability",
		JSONAttributesTopic: "homeassistant/sensor/weather_station_temperature/attributes",
		Device: DeviceInfo{
			Identifiers:  []string{cfg.DeviceID},
			Name:         cfg.DeviceName,
			Manufacturer: cfg.DeviceManufacturer,
			Model:        cfg.DeviceModel,
		},
		Origin: OriginInfo{
			Name:       "VEVOR Weatherbridge",
			SWVersion:  "0.1.0",
			SupportURL: "https://github.com/lenucksi/VevorWeatherbridge",
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	// Unmarshal to check structure
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	// Check required fields
	if result["name"] != "Weather Station Temperature" {
		t.Errorf("name = %v, want Weather Station Temperature", result["name"])
	}
	if result["device_class"] != "temperature" {
		t.Errorf("device_class = %v, want temperature", result["device_class"])
	}
	if result["unit_of_measurement"] != "°C" {
		t.Errorf("unit_of_measurement = %v, want °C", result["unit_of_measurement"])
	}

	// Check device info
	device, ok := result["device"].(map[string]interface{})
	if !ok {
		t.Fatal("device field missing or wrong type")
	}
	if device["manufacturer"] != "VEVOR" {
		t.Errorf("device.manufacturer = %v, want VEVOR", device["manufacturer"])
	}

	// Check origin info
	origin, ok := result["origin"].(map[string]interface{})
	if !ok {
		t.Fatal("origin field missing or wrong type")
	}
	if origin["name"] != "VEVOR Weatherbridge" {
		t.Errorf("origin.name = %v, want VEVOR Weatherbridge", origin["name"])
	}
}

func TestDiscoveryPayloadOmitsEmptyFields(t *testing.T) {
	payload := DiscoveryPayload{
		Name:              "Weather Station UV Index",
		StateTopic:        "homeassistant/sensor/weather_station_uv_index/state",
		UniqueID:          "weather_station_uv_index",
		UnitOfMeasurement: "index",
		AvailabilityTopic: "homeassistant/sensor/weather_station/availability",
		Device: DeviceInfo{
			Identifiers: []string{"weather_station"},
			Name:        "Weather Station",
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	// These fields should be omitted due to omitempty
	if _, ok := result["device_class"]; ok {
		t.Error("device_class should be omitted when empty")
	}
	if _, ok := result["icon"]; ok {
		t.Error("icon should be omitted when empty")
	}
	if _, ok := result["state_class"]; ok {
		t.Error("state_class should be omitted when empty")
	}
}

func TestMQTTClientTopics(t *testing.T) {
	cfg := &Config{
		MQTTPrefix: "homeassistant",
		DeviceID:   "weather_station",
	}

	// Create a mock client struct just to test topic generation
	m := &MQTTClient{cfg: cfg}

	tests := []struct {
		name     string
		fn       func(string) string
		sensorID string
		expected string
	}{
		{"ConfigTopic", m.ConfigTopic, "temperature", "homeassistant/sensor/weather_station_temperature/config"},
		{"StateTopic", m.StateTopic, "temperature", "homeassistant/sensor/weather_station_temperature/state"},
		{"AttributesTopic", m.AttributesTopic, "humidity", "homeassistant/sensor/weather_station_humidity/attributes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.sensorID)
			if result != tt.expected {
				t.Errorf("%s(%q) = %q, want %q", tt.name, tt.sensorID, result, tt.expected)
			}
		})
	}
}

func TestMQTTClientAvailabilityTopic(t *testing.T) {
	cfg := &Config{
		MQTTPrefix: "homeassistant",
		DeviceID:   "my_weather",
	}

	m := &MQTTClient{cfg: cfg}
	expected := "homeassistant/sensor/my_weather/availability"
	result := m.AvailabilityTopic()

	if result != expected {
		t.Errorf("AvailabilityTopic() = %q, want %q", result, expected)
	}
}

func TestConfigTimezoneLoading(t *testing.T) {
	// This tests that timezone loading works
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("Failed to load timezone: %v", err)
	}

	if loc.String() != "America/New_York" {
		t.Errorf("Timezone = %v, want America/New_York", loc.String())
	}
}
