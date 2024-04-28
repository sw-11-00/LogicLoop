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
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			name: "Test2",
			head: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
			want: &ListNode{2, &ListNode{3, nil}},
		},
		{
			name: "Test3",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
		{
			name: "Test4",
			head: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}},
			want: nil,
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

	dummy := ListNode{Val: 0, Next: head}
	prev := &dummy

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}

	return dummy.Next
}
