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

// FToC converts Fahrenheit to Celsius, rounded to 1 decimal place.
func FToC(f float64) float64 {
	return roundTo((f-32)*5.0/9.0, 1)
}

// InHgToHPa converts inches of mercury to hectopascals, rounded to 1 decimal place.
func InHgToHPa(inhg float64) float64 {
	return roundTo(inhg*33.8639, 1)
}

// MphToKmh converts miles per hour to kilometers per hour, rounded to 1 decimal place.
func MphToKmh(mph float64) float64 {
	return roundTo(mph*1.60934, 1)
}

// InchToMm converts inches to millimeters, rounded to 1 decimal place.
func InchToMm(inch float64) float64 {
	return roundTo(inch*25.4, 1)
}

// roundTo rounds a float64 to the specified number of decimal places.
func roundTo(val float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(val*pow) / pow
}
