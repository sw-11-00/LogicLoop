package LinkedList

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{"empty list", nil, nil},
		{"single node", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"two nodes", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 2}},
		{"three nodes", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, &ListNode{Val: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := middleNode(tt.head)
			for result != nil && tt.expected != nil {
				if result.Val != tt.expected.Val {
					t.Errorf("got %v, want %v", result.Val, tt.expected.Val)
				}
				result = result.Next
				tt.expected = tt.expected.Next
			}
			if (result != nil && tt.expected == nil) || (result == nil && tt.expected != nil) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
