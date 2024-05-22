package LinkedList

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected bool
	}{
		{"empty list", nil, true},
		{"single node", &ListNode{Val: 1}, true},
		{"multiple nodes, not palindrome", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		}, false},
		{"multiple nodes, palindrome", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.head)
			if result != tt.expected {
				t.Errorf("isPalindrome() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}

	for head.Next != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}

	return true
}
