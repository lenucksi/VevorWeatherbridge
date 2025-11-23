package main

import "math"

// cardinalDirections contains the 16-point compass rose directions.
var cardinalDirections = []string{
	"N", "NNE", "NE", "ENE",
	"E", "ESE", "SE", "SSE",
	"S", "SSW", "SW", "WSW",
	"W", "WNW", "NW", "NNW",
}

// DegreesToCardinal converts degrees (0-360) to a 16-point cardinal direction.
// Each direction covers 22.5 degrees, with North centered at 0 degrees.
// Returns empty string for NaN input.
func DegreesToCardinal(degrees float64) string {
	if math.IsNaN(degrees) {
		return ""
	}

	// Normalize to 0-360 range
	degrees = math.Mod(degrees, 360)
	if degrees < 0 {
		degrees += 360
	}

	// 11.25 offset centers N at 0; 22.5 degrees per segment
	index := int((degrees+11.25)/22.5) % 16
	return cardinalDirections[index]
}
