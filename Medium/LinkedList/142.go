package LinkedList

import (
	"fmt"
)

func DetectCycleTest() {
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

	// 测试
	testCases := []struct {
		head     *ListNode
		expected *ListNode
		desc     string
	}{
		{node1, node3, "List with cycle starting at node with value 3"},
		{node6, nil, "List without cycle"},
		{nil, nil, "Empty list"},
	}

	for _, tc := range testCases {
		result := detectCycle(tc.head)
		if (result == nil && tc.expected == nil) || (result != nil && result == tc.expected) {
			fmt.Printf("PASS: %s\n", tc.desc)
		} else {
			fmt.Printf("FAIL: %s\n", tc.desc)
		}
	}
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow := head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}
