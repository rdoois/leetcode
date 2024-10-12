package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, carry int
	res := &ListNode{}
	current := res

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry
		if l1 != nil {
			sum = l1.Val + sum
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 9 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	return res.Next
}
