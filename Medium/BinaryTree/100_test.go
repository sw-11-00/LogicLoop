package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		name     string
		tree1    *common.TreeNode
		tree2    *common.TreeNode
		expected bool
	}{
		{
			name:     "both empty",
			tree1:    nil,
			tree2:    nil,
			expected: true,
		},
		{
			name:     "one empty, one non-empty",
			tree1:    nil,
			tree2:    &common.TreeNode{Val: 1},
			expected: false,
		},
		{
			name: "both non-empty and same",
			tree1: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
				},
				Right: &common.TreeNode{
					Val: 3,
				},
			},
			tree2: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
				},
				Right: &common.TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			name: "both non-empty and different",
			tree1: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
				},
				Right: &common.TreeNode{
					Val: 3,
				},
			},
			tree2: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
				},
				Right: &common.TreeNode{
					Val: 4,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSameTree(tc.tree1, tc.tree2)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %t, got %t", tc.name, tc.expected, actual)
			}
		})
	}
}

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
