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
	"fmt"
	"log/slog"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	// Version is the application version for origin info.
	Version = "0.1.0"
	// SupportURL is the project URL for origin info.
	SupportURL = "https://github.com/lenucksi/VevorWeatherbridge"
)

// DiscoveryPayload represents a Home Assistant MQTT Discovery config message.
type DiscoveryPayload struct {
	Name                      string     `json:"name"`
	StateTopic                string     `json:"state_topic"`
	UniqueID                  string     `json:"unique_id"`
	DeviceClass               string     `json:"device_class,omitempty"`
	UnitOfMeasurement         string     `json:"unit_of_measurement,omitempty"`
	StateClass                string     `json:"state_class,omitempty"`
	Icon                      string     `json:"icon,omitempty"`
	SuggestedDisplayPrecision int        `json:"suggested_display_precision,omitempty"`
	Device                    DeviceInfo `json:"device"`
	AvailabilityTopic         string     `json:"availability_topic"`
	JSONAttributesTopic       string     `json:"json_attributes_topic,omitempty"`
	Origin                    OriginInfo `json:"origin,omitempty"`
}

// DeviceInfo represents device information for Home Assistant.
type DeviceInfo struct {
	Identifiers  []string `json:"identifiers"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"`
	Model        string   `json:"model"`
}

// OriginInfo represents the origin of the integration.
type OriginInfo struct {
	Name       string `json:"name"`
	SWVersion  string `json:"sw_version"`
	SupportURL string `json:"support_url"`
}

// MQTTClient wraps the paho MQTT client with application-specific logic.
type MQTTClient struct {
	client    mqtt.Client
	cfg       *Config
	connected bool
	mu        sync.RWMutex
}

// NewMQTTClient creates and connects a new MQTT client.
func NewMQTTClient(cfg *Config) (*MQTTClient, error) {
	m := &MQTTClient{cfg: cfg}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.MQTTHost, cfg.MQTTPort))
	opts.SetClientID(fmt.Sprintf("vevor-weatherbridge-%s", cfg.DeviceID))
	opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(10 * time.Second)

	// Set credentials if provided
	if cfg.MQTTUser != "" {
		opts.SetUsername(cfg.MQTTUser)
		opts.SetPassword(cfg.MQTTPassword)
	}

	// Set Last Will and Testament
	availabilityTopic := m.AvailabilityTopic()
	opts.SetWill(availabilityTopic, "offline", 1, true)

	// Set callbacks
	opts.SetOnConnectHandler(m.onConnect)
	opts.SetConnectionLostHandler(m.onConnectionLost)

	m.client = mqtt.NewClient(opts)

	// Connect
	token := m.client.Connect()
	if token.WaitTimeout(10 * time.Second) {
		if token.Error() != nil {
			return nil, fmt.Errorf("MQTT connection failed: %w", token.Error())
		}
	} else {
		slog.Warn("MQTT connection timeout, will retry in background")
	}

	return m, nil
}

// onConnect is called when the client connects to the broker.
func (m *MQTTClient) onConnect(client mqtt.Client) {
	m.mu.Lock()
	m.connected = true
	m.mu.Unlock()

	slog.Info("MQTT connected", "host", m.cfg.MQTTHost, "port", m.cfg.MQTTPort)

	// Publish online status
	availTopic := m.AvailabilityTopic()
	token := client.Publish(availTopic, 1, true, "online")
	token.Wait()
	if token.Error() != nil {
		slog.Error("Failed to publish availability status", "topic", availTopic, "error", token.Error())
	} else {
		slog.Debug("Published availability status", "topic", availTopic, "status", "online")
	}
}

// onConnectionLost is called when the connection is lost.
func (m *MQTTClient) onConnectionLost(client mqtt.Client, err error) {
	m.mu.Lock()
	m.connected = false
	m.mu.Unlock()

	slog.Warn("MQTT connection lost", "error", err)
}

// IsConnected returns true if the client is connected.
func (m *MQTTClient) IsConnected() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.connected
}

// AvailabilityTopic returns the availability topic for this device.
func (m *MQTTClient) AvailabilityTopic() string {
	return fmt.Sprintf("%s/sensor/%s/availability", m.cfg.MQTTPrefix, m.cfg.DeviceID)
}

// ConfigTopic returns the config topic for a sensor.
func (m *MQTTClient) ConfigTopic(sensorID string) string {
	return fmt.Sprintf("%s/sensor/%s_%s/config", m.cfg.MQTTPrefix, m.cfg.DeviceID, sensorID)
}

// StateTopic returns the state topic for a sensor.
func (m *MQTTClient) StateTopic(sensorID string) string {
	return fmt.Sprintf("%s/sensor/%s_%s/state", m.cfg.MQTTPrefix, m.cfg.DeviceID, sensorID)
}

// AttributesTopic returns the attributes topic for a sensor.
func (m *MQTTClient) AttributesTopic(sensorID string) string {
	return fmt.Sprintf("%s/sensor/%s_%s/attributes", m.cfg.MQTTPrefix, m.cfg.DeviceID, sensorID)
}

// PublishSensorConfig publishes the discovery config for a sensor.
func (m *MQTTClient) PublishSensorConfig(sensor *SensorDefinition) error {
	payload := DiscoveryPayload{
		Name:                fmt.Sprintf("%s %s", m.cfg.DeviceName, sensor.Name),
		StateTopic:          m.StateTopic(sensor.ID),
		UniqueID:            fmt.Sprintf("%s_%s", m.cfg.DeviceID, sensor.ID),
		UnitOfMeasurement:   sensor.GetUnit(m.cfg.IsMetric()),
		AvailabilityTopic:   m.AvailabilityTopic(),
		JSONAttributesTopic: m.AttributesTopic(sensor.ID),
		Device: DeviceInfo{
			Identifiers:  []string{m.cfg.DeviceID},
			Name:         m.cfg.DeviceName,
			Manufacturer: m.cfg.DeviceManufacturer,
			Model:        m.cfg.DeviceModel,
		},
		Origin: OriginInfo{
			Name:       "VEVOR Weatherbridge",
			SWVersion:  Version,
			SupportURL: SupportURL,
		},
	}

	// Set device class if defined
	if sensor.DeviceClass != nil {
		payload.DeviceClass = *sensor.DeviceClass
	}

	// Set icon if defined (only for sensors without device_class)
	if sensor.Icon != "" {
		payload.Icon = sensor.Icon
	}

	// Set precision if defined
	if sensor.Precision > 0 {
		payload.SuggestedDisplayPrecision = sensor.Precision
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal config payload: %w", err)
	}

	topic := m.ConfigTopic(sensor.ID)
	token := m.client.Publish(topic, 1, true, data)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish config: %w", token.Error())
	}

	slog.Debug("Published sensor config", "sensor", sensor.ID, "topic", topic)
	return nil
}

// PublishSensorState publishes the state value for a sensor.
func (m *MQTTClient) PublishSensorState(sensorID string, value string) error {
	topic := m.StateTopic(sensorID)
	token := m.client.Publish(topic, 1, true, value)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish state: %w", token.Error())
	}

	slog.Debug("Published sensor state", "sensor", sensorID, "value", value)
	return nil
}

// PublishSensorAttributes publishes the attributes for a sensor.
func (m *MQTTClient) PublishSensorAttributes(sensorID string, attrs map[string]interface{}) error {
	data, err := json.Marshal(attrs)
	if err != nil {
		return fmt.Errorf("failed to marshal attributes: %w", err)
	}

	topic := m.AttributesTopic(sensorID)
	token := m.client.Publish(topic, 1, true, data)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish attributes: %w", token.Error())
	}

	slog.Debug("Published sensor attributes", "sensor", sensorID)
	return nil
}

// Close disconnects the MQTT client gracefully.
func (m *MQTTClient) Close() {
	// Publish offline status before disconnecting
	availTopic := m.AvailabilityTopic()
	token := m.client.Publish(availTopic, 1, true, "offline")
	if token.WaitTimeout(2 * time.Second) {
		if token.Error() != nil {
			slog.Error("Failed to publish offline status", "topic", availTopic, "error", token.Error())
		} else {
			slog.Debug("Published availability status", "topic", availTopic, "status", "offline")
		}
	} else {
		slog.Warn("Timeout publishing offline status", "topic", availTopic)
	}

	m.client.Disconnect(1000)
	slog.Info("MQTT disconnected")
}
