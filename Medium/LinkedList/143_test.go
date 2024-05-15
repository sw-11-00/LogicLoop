package LinkedList

import (
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{"empty list", nil, nil},
		{"single node", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"multiple nodes", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.head)
			for tt.head != nil && tt.expected != nil {
				if tt.head.Val != tt.expected.Val {
					t.Errorf("got %v, want %v", tt.head.Val, tt.expected.Val)
				}
				tt.head = tt.head.Next
				tt.expected = tt.expected.Next
			}
		})
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the list
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Merge the two halves
	first, second := head, prev
	for second.Next != nil {
		firstNext := first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		first = firstNext
		second = secondNext
	}
}
