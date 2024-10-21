package main

import "testing"

func TestMyAtoi(t *testing.T) {
	tableTests := []struct {
		description string
		input       string
		expected    int
	}{
		{"Case 1", "42", 42},
		{"Case 2", " -042", -42},
		{"Case 3", "1337c0d3", 1337},
		{"Case 4", "0-1", 0},
		{"Case 5", "words and 987", 0},
		{"Case 6", "-+12", 0},
		{"Case 7", "+-12", 0},
		{"Case 8", "   -042", -42},
		{"Case 9", "9223372036854775808", 2147483647},
		{"Case 10", "-91283472332", -2147483648},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := myAtoi(tt.input)
			if res != tt.expected {
				t.Errorf("given %s got %d, want %d", tt.input, res, tt.expected)
			}
		})
	}

}
