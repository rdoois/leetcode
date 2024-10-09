package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tableTests := []struct {
		description string
		nums        []int
		target      int
		expected    []int
	}{
		{"Case 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Case 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Case 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
