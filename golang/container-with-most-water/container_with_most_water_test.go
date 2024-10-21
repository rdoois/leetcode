package main

import "testing"

func TestMaxArea(t *testing.T) {
	tableTests := []struct {
		description string
		input       []int
		expected    int
	}{
		{"Case 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Case 2", []int{1, 1}, 1},
		{"Case 3", []int{1, 2, 4, 3}, 4},
		{"Case 4", []int{8, 7, 2, 1}, 7},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := maxArea(tt.input)
			if res != tt.expected {
				t.Errorf("given %v got %d, want %d", tt.input, res, tt.expected)
			}
		})
	}

}
