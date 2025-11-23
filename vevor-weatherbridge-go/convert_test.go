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
