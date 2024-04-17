package LinkedList

import (
	"fmt"
)

func HasCycleTest() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	// 创建测试用例 2: 有环链表
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node4.Next = node5
	node5.Next = node6
	node6.Next = node5 // 创建环

	// 测试
	testCases := []struct {
		head     *ListNode
		expected bool
	}{
		{node1, false},
		{node4, true},
		{nil, false}, // 空链表测试
	}

	for i, tc := range testCases {
		result := hasCycle(tc.head)
		fmt.Printf("Test Case %d: Expected %t, Got %t\n", i, tc.expected, result)
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
