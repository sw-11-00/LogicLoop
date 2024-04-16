package LinkedList

import "fmt"

func ReverseLinkedListTest() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	reversedHead := reverseList(head)

	current := reversedHead
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	return prev
}
