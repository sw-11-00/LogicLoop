package LinkedList

import (
	"testing"
)

func TestKthToLast(t *testing.T) {
	// head = [1,2,3,4,5], k = 2
	// Output: 4
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	k := 2
	expected := 4
	got := kthToLast(head, k)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	// head = [1], k = 1
	// Output: 1
	head = &ListNode{Val: 1}
	k = 1
	expected = 1
	got = kthToLast(head, k)
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func kthToLast(head *ListNode, k int) int {
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Val
}
