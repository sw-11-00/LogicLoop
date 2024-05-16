package LinkedList

import (
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = nil

	got := removeDuplicateNodes(head)
	want := &ListNode{Val: 1}
	want.Next = &ListNode{Val: 2}
	want.Next.Next = &ListNode{Val: 3}
	want.Next.Next.Next = nil

	if !got.Equal(want) {
		t.Errorf("got %v want %v given", got, want)
	}
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	visited := make(map[int]struct{})
	var prev *ListNode
	current := head
	for current != nil {
		if _, existed := visited[current.Val]; existed {
			prev.Next = current.Next
		} else {
			visited[current.Val] = struct{}{}
			prev = current
		}
		current = current.Next
	}

	return head
}
