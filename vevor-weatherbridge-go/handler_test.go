package main

import (
	"testing"
)

func TestNormalizeTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single-digit seconds",
			input:    "2025-11-28 18:58:7",
			expected: "2025-11-28 18:58:07",
		},
		{
			name:     "single-digit minutes",
			input:    "2025-11-28 18:5:07",
			expected: "2025-11-28 18:05:07",
		},
		{
			name:     "single-digit hours",
			input:    "2025-11-28 1:58:07",
			expected: "2025-11-28 01:58:07",
		},
		{
			name:     "all single-digit time",
			input:    "2025-11-28 1:5:7",
			expected: "2025-11-28 01:05:07",
		},
		{
			name:     "already properly formatted",
			input:    "2025-11-28 18:58:07",
			expected: "2025-11-28 18:58:07",
		},
		{
			name:     "midnight with single digits",
			input:    "2025-11-28 0:0:0",
			expected: "2025-11-28 00:00:00",
		},
		{
			name:     "noon properly formatted",
			input:    "2025-11-28 12:00:00",
			expected: "2025-11-28 12:00:00",
		},
		{
			name:     "invalid format returns as-is",
			input:    "invalid timestamp",
			expected: "invalid timestamp",
		},
		{
			name:     "empty string returns as-is",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeTimestamp(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeTimestamp(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
