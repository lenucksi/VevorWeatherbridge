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
