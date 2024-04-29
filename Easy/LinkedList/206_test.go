package LinkedList

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Test1",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			want: &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
		{
			name: "Test2",
			head: &ListNode{1, &ListNode{2, nil}},
			want: &ListNode{2, &ListNode{1, nil}},
		},
		{
			name: "Test3",
			head: &ListNode{1, nil},
			want: &ListNode{1, nil},
		},
		{
			name: "Test4",
			head: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !compareLinkedList(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// prev := &ListNode{}这个为什么不行
	var prev *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	return prev
}

func compareLinkedList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
