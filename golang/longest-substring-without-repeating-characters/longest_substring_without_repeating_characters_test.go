package main

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tableTests := []struct {
		description string
		input       string
		expected    int
	}{
		{"Case 1", "abcabcbb", 3},
		{"Case 2", "bbbbb", 1},
		{"Case 3", "pwwkew", 3},
		{"Case 4", " ", 1},
		{"Case 5", "aab", 2},
		{"Case 6", "dvdf", 3},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := lengthOfLongestSubstring(tt.input)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
