package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tableTests := []struct {
		description string
		input       []string
		expected    string
	}{
		{"Case 1", []string{"flower", "flow", "flight"}, "fl"},
		{"Case 2", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := longestCommonPrefix(tt.input)
			if res != tt.expected {
				t.Errorf("given %v got %s, want %s", tt.input, res, tt.expected)
			}
		})
	}

}
