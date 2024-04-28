package LinkedList

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "Test1",
			list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			want:  &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			name:  "Test2",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "Test3",
			list1: nil,
			list2: &ListNode{0, nil},
			want:  &ListNode{0, nil},
		},
		{
			name:  "Test4",
			list1: &ListNode{-1, &ListNode{2, &ListNode{4, nil}}},
			list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			want:  &ListNode{-1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			name:  "Test5",
			list1: &ListNode{1000000, &ListNode{2000000, &ListNode{4000000, nil}}},
			list2: &ListNode{1000000, &ListNode{3000000, &ListNode{4000000, nil}}},
			want:  &ListNode{1000000, &ListNode{1000000, &ListNode{2000000, &ListNode{3000000, &ListNode{4000000, &ListNode{4000000, nil}}}}}},
		},
		{
			name:  "Test6",
			list1: &ListNode{5, &ListNode{9, &ListNode{11, nil}}},
			list2: &ListNode{4, &ListNode{8, &ListNode{15, nil}}},
			want:  &ListNode{4, &ListNode{5, &ListNode{8, &ListNode{9, &ListNode{11, &ListNode{15, nil}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !isEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{}
	current := &head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	for list1 != nil {
		current.Next = list1
		list1 = list1.Next
		current = current.Next
	}

	for list2 != nil {
		current.Next = list2
		list2 = list2.Next
		current = current.Next
	}

	return head.Next
}
