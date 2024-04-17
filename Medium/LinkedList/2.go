package LinkedList

import (
	"fmt"
)

func AddTwoNumbersTest() {
	// 创建测试用例
	// 用例1: (2 -> 4 -> 3) + (5 -> 6 -> 4) = (7 -> 0 -> 8)
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// 用例2: (0) + (0) = (0)
	l3 := &ListNode{0, nil}
	l4 := &ListNode{0, nil}

	// 用例3: (9 -> 9 -> 9 -> 9 -> 9) + (1) = (0 -> 0 -> 0 -> 0 -> 0 -> 1)
	l5 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}
	l6 := &ListNode{1, nil}

	// 测试并打印结果
	printList(addTwoNumbers(l1, l2)) // 应输出 7 -> 0 -> 8
	printList(addTwoNumbers(l3, l4)) // 应输出 0
	printList(addTwoNumbers(l5, l6)) // 应输出 0 -> 0 -> 0 -> 0 -> 0 -> 1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carry
		volume := sum % 10
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: volume}
			tail = head
		} else {
			tail.Next = &ListNode{Val: volume}
			tail = tail.Next
		}
	}

	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}

	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}
