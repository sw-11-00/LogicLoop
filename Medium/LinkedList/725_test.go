package LinkedList

import (
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want []*ListNode
	}{
		{
			name: "Test1",
			head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			k:    5,
			want: []*ListNode{&ListNode{1, nil}, &ListNode{2, nil}, &ListNode{3, nil}, nil, nil},
		},
		{
			name: "Test2",
			head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			k:    3,
			want: []*ListNode{&ListNode{1, nil}, &ListNode{2, nil}, &ListNode{3, nil}},
		},
		{
			name: "Test3",
			head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			k:    1,
			want: []*ListNode{&ListNode{1, nil}},
		},
		{
			name: "Test4",
			head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			k:    0,
			want: []*ListNode{&ListNode{1, nil}, &ListNode{2, nil}, &ListNode{3, nil}},
		},
		{
			name: "Test5",
			head: nil,
			k:    3,
			want: []*ListNode{nil, nil, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitListToParts(tt.head, tt.k)
			if !isEqualList(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqualList(got []*ListNode, want []*ListNode) bool {
	if len(got) != len(want) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if !isEqual(got[i], want[i]) {
			return false
		}
	}
	return true
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	return nil
}
