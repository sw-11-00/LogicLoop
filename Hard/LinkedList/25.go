package LinkedList

import "fmt"

func ReverseKGroupTest() {
	// test case 0
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	k := 3

	result := reverseKGroup(head, k)

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
