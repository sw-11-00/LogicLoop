package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

func TestMinDepth(t *testing.T) {
	testCases := []struct {
		name     string
		root     *common.TreeNode
		expected int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "single node",
			root:     &common.TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "one-sided left",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
					Left: &common.TreeNode{
						Val: 3,
					},
				},
			},
			expected: 3,
		},
		{
			name: "one-sided right",
			root: &common.TreeNode{
				Val: 1,
				Right: &common.TreeNode{
					Val: 2,
					Right: &common.TreeNode{
						Val: 3,
					},
				},
			},
			expected: 3,
		},
		{
			name: "unbalanced right heavier",
			root: &common.TreeNode{
				Val:  1,
				Left: &common.TreeNode{Val: 2},
				Right: &common.TreeNode{
					Val: 3,
					Right: &common.TreeNode{
						Val: 4,
					},
				},
			},
			expected: 2,
		},
		{
			name: "balanced tree",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val:   2,
					Left:  &common.TreeNode{Val: 4},
					Right: &common.TreeNode{Val: 5},
				},
				Right: &common.TreeNode{
					Val:   3,
					Left:  &common.TreeNode{Val: 6},
					Right: &common.TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := minDepth(tc.root)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return leftDepth + rightDepth + 1
	}

	return min(leftDepth, rightDepth) + 1
}
