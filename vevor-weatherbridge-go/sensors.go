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

// SensorDefinition contains metadata for a weather sensor.
type SensorDefinition struct {
	Name         string  // Human-readable name (e.g., "Temperature")
	ID           string  // Snake_case identifier (e.g., "temperature")
	QueryParam   string  // Weather Underground query parameter
	DeviceClass  *string // Home Assistant device class (nil if none)
	MetricUnit   string  // Unit in metric system
	ImperialUnit string  // Unit in imperial system
	Icon         string  // Material Design Icon (mdi:xxx), empty if device_class provides one
	Precision    int     // Suggested display precision (0 = not set)
}

// Helper to create a string pointer
func strPtr(s string) *string {
	return &s
}

// SensorDefinitions contains all supported weather sensors.
var SensorDefinitions = []SensorDefinition{
	{
		Name:         "Barometric Pressure",
		ID:           "barometric_pressure",
		QueryParam:   "baromin",
		DeviceClass:  strPtr("atmospheric_pressure"),
		MetricUnit:   "hPa",
		ImperialUnit: "inHg",
	},
	{
		Name:         "Temperature",
		ID:           "temperature",
		QueryParam:   "tempf",
		DeviceClass:  strPtr("temperature"),
		MetricUnit:   "°C",
		ImperialUnit: "°F",
	},
	{
		Name:         "Humidity",
		ID:           "humidity",
		QueryParam:   "humidity",
		DeviceClass:  strPtr("humidity"),
		MetricUnit:   "%",
		ImperialUnit: "%",
	},
	{
		Name:         "Dew Point",
		ID:           "dew_point",
		QueryParam:   "dewptf",
		DeviceClass:  strPtr("temperature"),
		MetricUnit:   "°C",
		ImperialUnit: "°F",
	},
	{
		Name:         "Rainfall",
		ID:           "rainfall",
		QueryParam:   "rainin",
		DeviceClass:  strPtr("precipitation"),
		MetricUnit:   "mm",
		ImperialUnit: "in",
	},
	{
		Name:         "Daily Rainfall",
		ID:           "daily_rainfall",
		QueryParam:   "dailyrainin",
		DeviceClass:  strPtr("precipitation"),
		MetricUnit:   "mm",
		ImperialUnit: "in",
	},
	{
		Name:         "Wind Direction",
		ID:           "wind_direction",
		QueryParam:   "winddir",
		DeviceClass:  nil,
		MetricUnit:   "°",
		ImperialUnit: "°",
		Icon:         "mdi:compass-outline",
		Precision:    0,
	},
	{
		Name:         "Wind Speed",
		ID:           "wind_speed",
		QueryParam:   "windspeedmph",
		DeviceClass:  strPtr("wind_speed"),
		MetricUnit:   "km/h",
		ImperialUnit: "mph",
	},
	{
		Name:         "Wind Gust Speed",
		ID:           "wind_gust_speed",
		QueryParam:   "windgustmph",
		DeviceClass:  strPtr("wind_speed"),
		MetricUnit:   "km/h",
		ImperialUnit: "mph",
	},
	{
		Name:         "UV Index",
		ID:           "uv_index",
		QueryParam:   "UV",
		DeviceClass:  nil,
		MetricUnit:   "index",
		ImperialUnit: "index",
	},
	{
		Name:         "Solar Radiation",
		ID:           "solar_radiation",
		QueryParam:   "solarRadiation",
		DeviceClass:  strPtr("irradiance"),
		MetricUnit:   "W/m²",
		ImperialUnit: "W/m²",
	},
}

// GetSensorByQueryParam returns the sensor definition for a query parameter.
func GetSensorByQueryParam(param string) *SensorDefinition {
	for i := range SensorDefinitions {
		if SensorDefinitions[i].QueryParam == param {
			return &SensorDefinitions[i]
		}
	}
	return nil
}

// GetUnit returns the appropriate unit based on metric/imperial setting.
func (s *SensorDefinition) GetUnit(isMetric bool) string {
	if isMetric {
		return s.MetricUnit
	}
	return s.ImperialUnit
}
