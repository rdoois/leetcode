package main

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	tableTests := []struct {
		description string
		brackets    string
		expected    bool
	}{
		{"Case 1", "()", true},
		{"Case 2", "()[]{}", true},
		{"Case 3", "(]", false},
		{"Case 4", "([])", true},
		{"Case 5", "([)]", false},
		{"Case 6", "[", false},
		{"Case 7", "((", false},
		{"Case 8", "){", false},
		{"Case 9", "()))", false},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := isValid(tt.brackets)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
