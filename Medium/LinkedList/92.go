package LinkedList

import "fmt"

func ReverseLinkedListII() {
	// test case 0
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	left := 2
	right := 4

	result := reverseBetween(head, left, right)

	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
	fmt.Println()

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return nil
}
