package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tableTests := []struct {
		description string
		l1          *ListNode
		l2          *ListNode
		expected    *ListNode
	}{
		{"Case 1", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}},
		{"Case 2", &ListNode{}, &ListNode{}, &ListNode{}},
		{"Case 3", &ListNode{}, &ListNode{Val: 0}, &ListNode{Val: 0}},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := mergeTwoLists(tt.l1, tt.l2)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}

}
