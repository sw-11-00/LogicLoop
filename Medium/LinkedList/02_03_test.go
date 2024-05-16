package LinkedList

import (
	"testing"
)

func TestDeleteNode1(t *testing.T) {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next = nil

	deleteNode1(head.Next)
	want := &ListNode{Val: 4}
	want.Next = &ListNode{Val: 1}
	want.Next.Next = &ListNode{Val: 9}
	want.Next.Next.Next = nil

	if !head.Equal(want) {
		t.Errorf("got %v want %v given", head, want)
	}
}

func deleteNode1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
