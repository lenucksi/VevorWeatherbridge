package main

import (
	"math"
	"testing"
)

func TestFToC(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"freezing point", 32.0, 0.0},
		{"boiling point", 212.0, 100.0},
		{"body temperature", 98.6, 37.0},
		{"negative", -40.0, -40.0}, // F and C are equal at -40
		{"room temperature", 68.0, 20.0},
		{"typical weather", 75.5, 24.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FToC(tt.input)
			if math.Abs(result-tt.expected) > 0.05 {
				t.Errorf("FToC(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestInHgToHPa(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"standard atmosphere", 29.92, 1013.2},
		{"low pressure", 29.0, 982.1},
		{"high pressure", 30.5, 1032.8},
		{"zero", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InHgToHPa(tt.input)
			if math.Abs(result-tt.expected) > 0.15 {
				t.Errorf("InHgToHPa(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMphToKmh(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"zero", 0.0, 0.0},
		{"walking speed", 3.0, 4.8},
		{"highway speed", 60.0, 96.6},
		{"wind gust", 25.5, 41.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MphToKmh(tt.input)
			if math.Abs(result-tt.expected) > 0.1 {
				t.Errorf("MphToKmh(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestInchToMm(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"zero", 0.0, 0.0},
		{"one inch", 1.0, 25.4},
		{"light rain", 0.1, 2.5},
		{"heavy rain", 2.5, 63.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InchToMm(tt.input)
			if math.Abs(result-tt.expected) > 0.1 {
				t.Errorf("InchToMm(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRoundTo(t *testing.T) {
	tests := []struct {
		name     string
		val      float64
		decimals int
		expected float64
	}{
		{"round down", 1.234, 1, 1.2},
		{"round up", 1.256, 1, 1.3},
		{"round half up", 1.25, 1, 1.3},
		{"two decimals", 1.2345, 2, 1.23},
		{"zero decimals", 1.6, 0, 2.0},
		{"negative", -1.234, 1, -1.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundTo(tt.val, tt.decimals)
			if result != tt.expected {
				t.Errorf("roundTo(%v, %v) = %v, want %v", tt.val, tt.decimals, result, tt.expected)
			}
		})
	}
}

// FuzzFToC fuzzes the Fahrenheit to Celsius conversion to find edge cases
// with NaN, Inf, and extreme values.
func FuzzFToC(f *testing.F) {
	// Seed corpus with typical weather values
	f.Add(32.0)    // Freezing point
	f.Add(212.0)   // Boiling point
	f.Add(0.0)     // Zero
	f.Add(-40.0)   // F and C equal point
	f.Add(98.6)    // Body temperature
	f.Add(-459.67) // Absolute zero in Fahrenheit

	f.Fuzz(func(t *testing.T, fahrenheit float64) {
		result := FToC(fahrenheit)

		// Should not return NaN for valid inputs
		if !math.IsNaN(fahrenheit) && math.IsNaN(result) {
			t.Errorf("FToC(%v) returned NaN for non-NaN input", fahrenheit)
		}

		// If input is finite, output should be finite
		if !math.IsInf(fahrenheit, 0) && math.IsInf(result, 0) {
			t.Errorf("FToC(%v) returned Inf for finite input", fahrenheit)
		}

		// Verify inverse relationship (within rounding tolerance)
		// Note: FToC rounds to 1 decimal place, so we need larger tolerance
		if !math.IsNaN(result) && !math.IsInf(result, 0) {
			// C = (F - 32) * 5/9, so F = C * 9/5 + 32
			inverse := result*9.0/5.0 + 32.0
			// Allow tolerance of 0.2F due to 0.1C rounding
			if math.Abs(inverse-fahrenheit) > 0.2 {
				t.Errorf("FToC(%v) = %v, but inverse gives %v", fahrenheit, result, inverse)
			}
		}
	})
}

// FuzzInHgToHPa fuzzes the inches of mercury to hectopascals conversion.
func FuzzInHgToHPa(f *testing.F) {
	// Seed corpus with typical atmospheric pressure values
	f.Add(29.92) // Standard atmosphere
	f.Add(29.0)  // Low pressure
	f.Add(30.5)  // High pressure
	f.Add(0.0)   // Zero

	f.Fuzz(func(t *testing.T, inhg float64) {
		result := InHgToHPa(inhg)

		// Should not return NaN for valid inputs
		if !math.IsNaN(inhg) && math.IsNaN(result) {
			t.Errorf("InHgToHPa(%v) returned NaN for non-NaN input", inhg)
		}

		// If input is finite, output should be finite
		if !math.IsInf(inhg, 0) && math.IsInf(result, 0) {
			t.Errorf("InHgToHPa(%v) returned Inf for finite input", inhg)
		}

		// Result should maintain sign (positive input -> positive output)
		if inhg > 0 && result <= 0 {
			t.Errorf("InHgToHPa(%v) = %v, sign mismatch", inhg, result)
		}
		if inhg < 0 && result >= 0 {
			t.Errorf("InHgToHPa(%v) = %v, sign mismatch", inhg, result)
		}
	})
}

// FuzzMphToKmh fuzzes the miles per hour to kilometers per hour conversion.
func FuzzMphToKmh(f *testing.F) {
	// Seed corpus with typical wind speed values
	f.Add(0.0)   // Calm
	f.Add(5.0)   // Light breeze
	f.Add(25.0)  // Strong wind
	f.Add(100.0) // Hurricane force

	f.Fuzz(func(t *testing.T, mph float64) {
		result := MphToKmh(mph)

		// Should not return NaN for valid inputs
		if !math.IsNaN(mph) && math.IsNaN(result) {
			t.Errorf("MphToKmh(%v) returned NaN for non-NaN input", mph)
		}

		// If input is finite, output should be finite
		if !math.IsInf(mph, 0) && math.IsInf(result, 0) {
			t.Errorf("MphToKmh(%v) returned Inf for finite input", mph)
		}

		// Result should be ~1.609 times input (within rounding)
		if !math.IsNaN(result) && !math.IsInf(result, 0) && mph != 0 {
			ratio := result / mph
			if math.Abs(ratio-1.60934) > 0.01 {
				t.Errorf("MphToKmh(%v) = %v, ratio %v != 1.60934", mph, result, ratio)
			}
		}
	})
}

// FuzzInchToMm fuzzes the inches to millimeters conversion.
func FuzzInchToMm(f *testing.F) {
	// Seed corpus with typical rainfall values
	f.Add(0.0) // No rain
	f.Add(0.1) // Light rain
	f.Add(1.0) // Moderate rain
	f.Add(5.0) // Heavy rain

	f.Fuzz(func(t *testing.T, inch float64) {
		result := InchToMm(inch)

		// Should not return NaN for valid inputs
		if !math.IsNaN(inch) && math.IsNaN(result) {
			t.Errorf("InchToMm(%v) returned NaN for non-NaN input", inch)
		}

		// If input is finite, output should be finite
		if !math.IsInf(inch, 0) && math.IsInf(result, 0) {
			t.Errorf("InchToMm(%v) returned Inf for finite input", inch)
		}

		// Result should be ~25.4 times input (within rounding tolerance)
		// Note: InchToMm rounds to 1 decimal place, so ratio can vary
		if !math.IsNaN(result) && !math.IsInf(result, 0) && inch != 0 {
			ratio := result / inch
			// Allow larger tolerance for small values due to rounding
			tolerance := 1.0
			if math.Abs(inch) > 1.0 {
				tolerance = 0.5
			}
			if math.Abs(ratio-25.4) > tolerance {
				t.Errorf("InchToMm(%v) = %v, ratio %v != 25.4", inch, result, ratio)
			}
		}
	})
}

// FuzzRoundTo fuzzes the rounding function with various precision levels.
func FuzzRoundTo(f *testing.F) {
	// Seed corpus with various values and precisions
	f.Add(1.234, 1)
	f.Add(1.256, 2)
	f.Add(1.5, 0)
	f.Add(-1.234, 1)
	f.Add(0.0, 5)

	f.Fuzz(func(t *testing.T, val float64, decimals int) {
		// Should not panic regardless of input
		result := roundTo(val, decimals)

		// If input is NaN, output should be NaN
		if math.IsNaN(val) && !math.IsNaN(result) {
			t.Errorf("roundTo(NaN, %d) = %v, expected NaN", decimals, result)
		}

		// If input is Inf, output should be Inf with same sign
		if math.IsInf(val, 1) && !math.IsInf(result, 1) {
			t.Errorf("roundTo(+Inf, %d) = %v, expected +Inf", decimals, result)
		}
		if math.IsInf(val, -1) && !math.IsInf(result, -1) {
			t.Errorf("roundTo(-Inf, %d) = %v, expected -Inf", decimals, result)
		}

		// For finite values, result should be close to input
		if !math.IsNaN(val) && !math.IsInf(val, 0) && !math.IsNaN(result) && !math.IsInf(result, 0) {
			diff := math.Abs(result - val)
			// Difference should not be larger than 1 (rounding can't change by more than 1 at any decimal place)
			if diff > 1.0 {
				t.Errorf("roundTo(%v, %d) = %v, difference %v > 1.0", val, decimals, result, diff)
			}
		}
	})
}
