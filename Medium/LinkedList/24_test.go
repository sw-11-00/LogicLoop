package LinkedList

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single node", []int{1}, []int{1}},
		{"two nodes", []int{1, 2}, []int{2, 1}},
		{"multiple nodes", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := swapPairs(head)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("got %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next
		first.Next = second.Next
		current.Next = second
		current.Next.Next = first
		current = current.Next.Next
	}
	return dummy.Next
}
