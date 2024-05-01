package LinkedList

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "Test1",
			l1:   CreateLinkedList([]int{2, 4, 3}, -1),
			l2:   CreateLinkedList([]int{5, 6, 4}, -1),
			want: CreateLinkedList([]int{7, 0, 8}, -1),
		},
		{
			name: "Test2",
			l1:   CreateLinkedList([]int{0}, -1),
			l2:   CreateLinkedList([]int{0}, -1),
			want: CreateLinkedList([]int{0}, -1),
		},
		{
			name: "Test3",
			l1:   CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9}, -1),
			l2:   CreateLinkedList([]int{9, 9, 9, 9}, -1),
			want: CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}, -1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers1(tt.l1, tt.l2); !compareLinkedList(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
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

func CreateLinkedList(nums []int, cyclePos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(nums))
	for i, num := range nums {
		nodes[i] = &ListNode{Val: num}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if cyclePos != -1 {
		nodes[len(nums)-1].Next = nodes[cyclePos]
	}

	return nodes[0]
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
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
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: sum % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum % 10}
			tail = tail.Next
		}
	}

	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}

	return head
}
