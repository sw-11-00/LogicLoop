package LinkedList

import (
	"testing"
)

func TestReverseLinkedListII(t *testing.T) {
	tests := []struct {
		name  string
		head  *ListNode
		left  int
		right int
		want  *ListNode
	}{
		{
			name:  "Test1",
			head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			left:  2,
			right: 4,
			want:  &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}},
		},
		{
			name:  "Test2",
			head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			left:  1,
			right: 5,
			want:  &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
		{
			name:  "Test3",
			head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			left:  3,
			right: 3,
			want:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name:  "Test4",
			head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			left:  1,
			right: 1,
			want:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name:  "Test5",
			head:  nil,
			left:  1,
			right: 1,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.head, tt.left, tt.right); !compareLinkedList(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n ListNode) Equal(want *ListNode) bool {
	return compareLinkedList(&n, want)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	p := dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	var prev *ListNode
	current := p.Next

	for i := 0; i < right-left+1; i++ {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	p.Next.Next = current
	p.Next = prev

	return dummy.Next
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
