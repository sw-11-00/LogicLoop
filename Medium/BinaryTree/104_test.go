package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func TestMaxDepth(t *testing.T) {
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
			name: "complete tree",
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
			actual := maxDepth(tc.root)
			if actual != tc.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
