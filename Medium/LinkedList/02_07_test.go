package LinkedList

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	headA := &ListNode{Val: 4}
	headA.Next = &ListNode{Val: 1}
	headA.Next.Next = &ListNode{Val: 8}
	headA.Next.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next.Next = &ListNode{Val: 5}
	headA.Next.Next.Next.Next.Next = nil

	headB := &ListNode{Val: 5}
	headB.Next = &ListNode{Val: 0}
	headB.Next.Next = &ListNode{Val: 1}
	headB.Next.Next.Next = &ListNode{Val: 8}
	headB.Next.Next.Next.Next = &ListNode{Val: 4}
	headB.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	headB.Next.Next.Next.Next.Next.Next = nil

	got := getIntersectionNode(headA, headB)
	want := &ListNode{Val: 8}
	want.Next = &ListNode{Val: 4}
	want.Next.Next = &ListNode{Val: 5}
	want.Next.Next.Next = nil

	if !got.Equal(want) {
		t.Errorf("got %v want %v given", got, want)
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
