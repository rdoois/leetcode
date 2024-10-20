package main

import "testing"

func TestReverse(t *testing.T) {
	tableTests := []struct {
		description string
		input       int
		expected    int
	}{
		{"Case 1", 123, 321},
		{"Case 2", -123, -321},
		{"Case 3", 120, 21},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := reverse(tt.input)
			if res != tt.expected {
				t.Errorf("given %d got %d, want %d", tt.input, res, tt.expected)
			}
		})
	}

}
