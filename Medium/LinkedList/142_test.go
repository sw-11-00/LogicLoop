package LinkedList

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pos  int
		want int
	}{
		{
			name: "Test1",
			nums: []int{3, 2, 0, -4},
			pos:  1,
			want: 2,
		},
		{
			name: "Test2",
			nums: []int{1, 2},
			pos:  0,
			want: 1,
		},
		{
			name: "Test3",
			nums: []int{1},
			pos:  -1,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := CreateLinkedList(tt.nums, tt.pos)
			got := detectCycle(head)
			if (got == nil && tt.want != -1) || (got != nil && got.Val != tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow
		}
	}

	return nil
}
