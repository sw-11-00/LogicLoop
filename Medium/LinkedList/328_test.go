package LinkedList

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "empty list",
			head: nil,
			want: nil,
		},
		{
			name: "single element",
			head: &ListNode{1, nil},
			want: &ListNode{1, nil},
		},
		{
			name: "two elements",
			head: &ListNode{1, &ListNode{2, nil}},
			want: &ListNode{1, &ListNode{2, nil}},
		},
		{
			name: "multiple elements",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			want: &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.head); !isEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func oddEvenList(head *ListNode) *ListNode {
	return nil
}
