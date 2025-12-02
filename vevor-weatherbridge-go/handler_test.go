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

// FuzzParseTimestamp fuzzes the parseTimestamp function to find edge cases
// in timestamp parsing from the weather station.
func FuzzParseTimestamp(f *testing.F) {
	// Seed corpus with known valid inputs from actual weather station
	f.Add("2025-12-1 11:15:31")   // Non-zero-padded day (actual device format)
	f.Add("2025-12-01 11:15:31")  // Zero-padded standard format
	f.Add("2025-1-5 1:5:7")       // All components non-padded
	f.Add("2025-11-28 18:58:7")   // Single-digit seconds
	f.Add("2025-11-28 0:0:0")     // Midnight
	f.Add("2025-12-31 23:59:59")  // End of year

	f.Fuzz(func(t *testing.T, timestamp string) {
		// Should not panic regardless of input
		result, err := parseTimestamp(timestamp)

		// If parsing succeeds, result should be valid (non-zero)
		if err == nil && result.IsZero() {
			t.Errorf("parseTimestamp(%q) returned zero time without error", timestamp)
		}

		// If parsing succeeds, we should be able to format it back
		if err == nil {
			formatted := result.Format("2006-01-02 15:04:05")
			if formatted == "" {
				t.Errorf("parseTimestamp(%q) succeeded but result cannot be formatted", timestamp)
			}
		}
	})
}
