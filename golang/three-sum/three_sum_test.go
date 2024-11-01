package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tableTests := []struct {
		description string
		nums        []int
		expected    [][]int
	}{
		{"Case 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"Case 2", []int{0, 1, 1}, [][]int{}},
		{"Case 3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := threeSum(tt.nums)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("given %v got %v, want %v", tt.nums, res, tt.expected)
			}
		})
	}

}
