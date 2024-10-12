package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tableTests := []struct {
		description string
		l1          *ListNode
		l2          *ListNode
		expected    *ListNode
	}{
		{"Case 1", &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}, &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}},
		{"Case 2", &ListNode{Val: 0}, &ListNode{Val: 0}, &ListNode{Val: 0}},
		{"Case 3", &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}, &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}}},
		{"Case 4", &ListNode{Val: 5, Next: &ListNode{Val: 6}}, &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}, &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}},
		{"Case 5", &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1}}}, &ListNode{Val: 1}, &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2}}}},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
