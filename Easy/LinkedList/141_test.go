package LinkedList

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node3 // 创建环，环起点为node3

	// 用例2：无环链表
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}
	node8 := &ListNode{Val: 8}

	node6.Next = node7
	node7.Next = node8

	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "Test1",
			head: node1,
			want: true,
		},
		{
			name: "Test2",
			head: node6,
			want: false,
		},
		{
			name: "Test3",
			head: nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
