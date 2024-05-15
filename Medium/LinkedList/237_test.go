package LinkedList

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		node     *ListNode
		expected *ListNode
	}{
		{"empty list", nil, &ListNode{Val: 1}, nil},
		{"single node", &ListNode{Val: 1}, &ListNode{Val: 1}, nil},
		{"multiple nodes", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}, &ListNode{Val: 3}, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.node)
			for tt.head != nil && tt.expected != nil {
				if tt.head.Val != tt.expected.Val {
					t.Errorf("got %v, want %v", tt.head.Val, tt.expected.Val)
				}
				tt.head = tt.head.Next
				tt.expected = tt.expected.Next
			}
		})
	}
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
