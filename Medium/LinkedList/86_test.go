package LinkedList

import (
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		x        int
		expected *ListNode
	}{
		{"empty list", nil, 3, nil},
		{"single node", &ListNode{Val: 1}, 2, &ListNode{Val: 1}},
		{"multiple nodes", &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		}, 3, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.head, tt.x)
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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	for fast.Next != nil {
		if fast.Next.Val < x {
			if slow == fast {
				slow = slow.Next
				fast = fast.Next
			} else {
				tmp := slow.Next
				slow.Next = fast.Next
				fast.Next = fast.Next.Next
				slow.Next.Next = tmp
				slow = slow.Next
			}
		} else {
			fast = fast.Next
		}
	}

	return dummy.Next
}
