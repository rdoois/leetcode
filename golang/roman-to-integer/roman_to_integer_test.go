package main

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tableTests := []struct {
		description string
		input       string
		expected    int
	}{
		{"Case 1", "III", 3},
		{"Case 2", "LVIII", 58},
		{"Case 3", "MCMXCIV", 1994},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := romanToInt(tt.input)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
