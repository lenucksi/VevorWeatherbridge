package main

import (
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldError bool
		expected    time.Time
	}{
		{
			name:        "non-zero-padded day (actual weather station format)",
			input:       "2025-12-1 11:15:31",
			shouldError: false,
			expected:    time.Date(2025, 12, 1, 11, 15, 31, 0, time.UTC),
		},
		{
			name:        "non-zero-padded month and day",
			input:       "2025-1-5 11:15:31",
			shouldError: false,
			expected:    time.Date(2025, 1, 5, 11, 15, 31, 0, time.UTC),
		},
		{
			name:        "zero-padded standard format",
			input:       "2025-12-01 11:15:31",
			shouldError: false,
			expected:    time.Date(2025, 12, 1, 11, 15, 31, 0, time.UTC),
		},
		{
			name:        "single-digit seconds",
			input:       "2025-11-28 18:58:7",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 18, 58, 7, 0, time.UTC),
		},
		{
			name:        "single-digit minutes",
			input:       "2025-11-28 18:5:07",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 18, 5, 7, 0, time.UTC),
		},
		{
			name:        "single-digit hours",
			input:       "2025-11-28 1:58:07",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 1, 58, 7, 0, time.UTC),
		},
		{
			name:        "all single-digit time components",
			input:       "2025-1-5 1:5:7",
			shouldError: false,
			expected:    time.Date(2025, 1, 5, 1, 5, 7, 0, time.UTC),
		},
		{
			name:        "already properly formatted",
			input:       "2025-11-28 18:58:07",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 18, 58, 7, 0, time.UTC),
		},
		{
			name:        "midnight with single digits",
			input:       "2025-11-28 0:0:0",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "noon properly formatted",
			input:       "2025-11-28 12:00:00",
			shouldError: false,
			expected:    time.Date(2025, 11, 28, 12, 0, 0, 0, time.UTC),
		},
		{
			name:        "invalid format returns error",
			input:       "invalid timestamp",
			shouldError: true,
		},
		{
			name:        "empty string returns error",
			input:       "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTimestamp(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("parseTimestamp(%q) expected error but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("parseTimestamp(%q) unexpected error: %v", tt.input, err)
				}
				if !result.Equal(tt.expected) {
					t.Errorf("parseTimestamp(%q) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
