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
