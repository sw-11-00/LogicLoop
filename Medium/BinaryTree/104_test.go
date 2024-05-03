package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

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

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
