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
