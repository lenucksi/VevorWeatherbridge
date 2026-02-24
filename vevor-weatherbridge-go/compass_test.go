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
	"math"
	"testing"
)

func TestDegreesToCardinal(t *testing.T) {
	tests := []struct {
		name     string
		degrees  float64
		expected string
	}{
		// Cardinal directions
		{"north at 0", 0, "N"},
		{"north at 360", 360, "N"},
		{"east at 90", 90, "E"},
		{"south at 180", 180, "S"},
		{"west at 270", 270, "W"},

		// Intercardinal directions
		{"northeast at 45", 45, "NE"},
		{"southeast at 135", 135, "SE"},
		{"southwest at 225", 225, "SW"},
		{"northwest at 315", 315, "NW"},

		// Secondary intercardinal directions
		{"north-northeast at 22.5", 22.5, "NNE"},
		{"east-northeast at 67.5", 67.5, "ENE"},
		{"east-southeast at 112.5", 112.5, "ESE"},
		{"south-southeast at 157.5", 157.5, "SSE"},
		{"south-southwest at 202.5", 202.5, "SSW"},
		{"west-southwest at 247.5", 247.5, "WSW"},
		{"west-northwest at 292.5", 292.5, "WNW"},
		{"north-northwest at 337.5", 337.5, "NNW"},

		// Edge cases - boundaries
		{"just before NNE boundary", 11.24, "N"},
		{"at NNE boundary", 11.25, "NNE"},
		{"just after NNE boundary", 11.26, "NNE"},
		{"just before N boundary from NNW", 348.74, "NNW"},
		{"at N boundary from NNW", 348.75, "N"},

		// Negative values (should normalize)
		{"negative -90 is west", -90, "W"},
		{"negative -180 is south", -180, "S"},
		{"negative -45 is NW", -45, "NW"},

		// Large values (should normalize)
		{"720 degrees is N", 720, "N"},
		{"450 degrees is E", 450, "E"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DegreesToCardinal(tt.degrees)
			if result != tt.expected {
				t.Errorf("DegreesToCardinal(%v) = %v, want %v", tt.degrees, result, tt.expected)
			}
		})
	}
}

func TestDegreesToCardinalNaN(t *testing.T) {
	result := DegreesToCardinal(math.NaN())
	if result != "" {
		t.Errorf("DegreesToCardinal(NaN) = %v, want empty string", result)
	}
}

func TestDegreesToCardinalAllDirections(t *testing.T) {
	// Test that all 16 directions are reachable
	expectedDirections := map[string]bool{
		"N": false, "NNE": false, "NE": false, "ENE": false,
		"E": false, "ESE": false, "SE": false, "SSE": false,
		"S": false, "SSW": false, "SW": false, "WSW": false,
		"W": false, "WNW": false, "NW": false, "NNW": false,
	}

	// Test at the center of each direction
	for i := 0; i < 16; i++ {
		degrees := float64(i) * 22.5
		direction := DegreesToCardinal(degrees)
		expectedDirections[direction] = true
	}

	for dir, found := range expectedDirections {
		if !found {
			t.Errorf("Direction %s was never returned", dir)
		}
	}
}

// FuzzDegreesToCardinal fuzzes the degrees to cardinal direction conversion
// to find edge cases with large numbers, negative values, and special floats.
func FuzzDegreesToCardinal(f *testing.F) {
	// Seed corpus with typical wind direction values
	f.Add(0.0)   // North
	f.Add(90.0)  // East
	f.Add(180.0) // South
	f.Add(270.0) // West
	f.Add(45.0)  // Northeast
	f.Add(22.5)  // North-northeast
	f.Add(360.0) // Full circle
	f.Add(-45.0) // Negative value

	f.Fuzz(func(t *testing.T, degrees float64) {
		// Should not panic regardless of input
		result := DegreesToCardinal(degrees)

		// NaN should return empty string
		if math.IsNaN(degrees) && result != "" {
			t.Errorf("DegreesToCardinal(NaN) = %q, expected empty string", result)
		}

		// Valid direction should be one of the 16 cardinal directions or empty (for NaN)
		validDirections := map[string]bool{
			"":  true, // Empty for NaN
			"N": true, "NNE": true, "NE": true, "ENE": true,
			"E": true, "ESE": true, "SE": true, "SSE": true,
			"S": true, "SSW": true, "SW": true, "WSW": true,
			"W": true, "WNW": true, "NW": true, "NNW": true,
		}

		if !validDirections[result] {
			t.Errorf("DegreesToCardinal(%v) = %q, not a valid direction", degrees, result)
		}

		// For non-NaN, non-Inf values, verify normalization works
		if !math.IsNaN(degrees) && !math.IsInf(degrees, 0) {
			// Normalize degrees to 0-360 range manually
			normalized := math.Mod(degrees, 360)
			if normalized < 0 {
				normalized += 360
			}

			// Result should be consistent with normalized value
			resultFromNormalized := DegreesToCardinal(normalized)
			if result != resultFromNormalized {
				t.Errorf("DegreesToCardinal(%v) = %q, but normalized DegreesToCardinal(%v) = %q",
					degrees, result, normalized, resultFromNormalized)
			}
		}
	})
}
