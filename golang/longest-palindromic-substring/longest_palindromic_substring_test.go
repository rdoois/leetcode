package main

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tableTests := []struct {
		description string
		input       string
		expected    string
	}{
		{"Case 1", "babad", "bab"},
		{"Case 2", "cbbd", "bb"},
		{"Case 3", "a", "a"},
		{"Case 4", "ac", "a"},
		{"Case 5", "ccc", "ccc"},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := longestPalindrome(tt.input)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
