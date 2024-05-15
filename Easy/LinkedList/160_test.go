package LinkedList

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name     string
		headA    *ListNode
		headB    *ListNode
		expected *ListNode
	}{
		{"both lists are empty", nil, nil, nil},
		{"no intersection", &ListNode{Val: 1}, &ListNode{Val: 2}, nil},
		{"has intersection", func() *ListNode {
			shared := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
			headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
			return headA
		}(), func() *ListNode {
			shared := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
			headB := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: shared}}}
			return headB
		}(), &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getIntersectionNode(tt.headA, tt.headB)
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
