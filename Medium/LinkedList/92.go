package LinkedList

import "fmt"

func ReverseLinkedListII() {
	// test case 0
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	left := 3
	right := 5

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
	dummy := &ListNode{Next: head}
	p := dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	var prev *ListNode
	current := p.Next

	for i := 0; i < right-left+1; i++ {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	p.Next.Next = current
	p.Next = prev

	return dummy.Next
}
