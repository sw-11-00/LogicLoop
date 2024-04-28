package LinkedList

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Test1",
			head: &ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			want: &ListNode{1, &ListNode{2, nil}},
		},
		{
			name: "Test2",
			head: &ListNode{1, &ListNode{1, &ListNode{1, nil}}},
			want: &ListNode{1, nil},
		},
		{
			name: "Test3",
			head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			name: "Test4",
			head: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}},
			want: &ListNode{1, &ListNode{2, nil}},
		},
		{
			name: "Test5",
			head: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			if !isEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
