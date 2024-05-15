package LinkedList

import (
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			name: "Test1",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    2,
			want: &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
		},
		{
			name: "Test2",
			head: &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			k:    4,
			want: &ListNode{2, &ListNode{0, &ListNode{1, nil}}},
		},
		{
			name: "Test3",
			head: &ListNode{1, &ListNode{2, nil}},
			k:    1,
			want: &ListNode{2, &ListNode{1, nil}},
		},
		{
			name: "Test4",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    5,
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "Test5",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    0,
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight1(tt.head, tt.k)
			if !isEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k %= length
	if k == 0 {
		return head
	}

	tail.Next = head
	for i := 0; i < length-k; i++ {
		tail = tail.Next
	}

	newHead := tail.Next
	tail.Next = nil

	return newHead
}
