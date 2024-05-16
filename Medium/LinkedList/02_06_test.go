package LinkedList

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next = nil

	got := isPalindrome(head)

	if got != true {
		t.Errorf("got %v want %v given", got, true)
	}
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var linkedArray []int
	for head != nil {
		linkedArray = append(linkedArray, head.Val)
		head = head.Next
	}

	for i := 0; i < len(linkedArray)/2; i++ {
		if linkedArray[i] != linkedArray[len(linkedArray)-i-1] {
			return false
		}
	}

	return true
}
