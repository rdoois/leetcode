package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tableTests := []struct {
		description string
		input       int
		expected    bool
	}{
		{"Case 1", 121, true},
		{"Case 2", -121, false},
		{"Case 3", 10, false},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := isPalindrome(tt.input)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
