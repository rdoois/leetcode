package main

import "testing"

func TestMaxArea(t *testing.T) {
	tableTests := []struct {
		description string
		input       int
		expected    string
	}{
		{"Case 1", 3749, "MMMDCCXLIX"},
		{"Case 2", 58, "LVIII"},
		{"Case 3", 1994, "MCMXCIV"},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := intToRoman(tt.input)
			if res != tt.expected {
				t.Errorf("given %d got %s, want %s", tt.input, res, tt.expected)
			}
		})
	}

}
