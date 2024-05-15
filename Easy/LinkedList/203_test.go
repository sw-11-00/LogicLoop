package LinkedList

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		val      int
		expected *ListNode
	}{
		{"empty list", nil, 1, nil},
		{"single node, value matches", &ListNode{Val: 1}, 1, nil},
		{"single node, value does not match", &ListNode{Val: 1}, 2, &ListNode{Val: 1}},
		{"multiple nodes, some values match", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		}, 1, &ListNode{
			Val: 2,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElements(tt.head, tt.val)
			for result != nil && tt.expected != nil {
				if result.Val != tt.expected.Val {
					t.Errorf("got %v, want %v", result.Val, tt.expected.Val)
				}
				result = result.Next
				tt.expected = tt.expected.Next
			}
			if (result != nil && tt.expected == nil) || (result == nil && tt.expected != nil) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev, current := dummy, head
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return dummy.Next
}
