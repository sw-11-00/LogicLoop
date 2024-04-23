package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		name     string
		root     *common.TreeNode
		expected bool
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "single node",
			root:     &common.TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "unbalanced right heavy",
			root: &common.TreeNode{
				Val: 1,
				Right: &common.TreeNode{
					Val: 2,
					Right: &common.TreeNode{
						Val: 3,
					},
				},
			},
			expected: false,
		},
		{
			name: "balanced tree",
			root: &common.TreeNode{
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isBalanced(tc.root)
			if actual != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

func isBalanced(root *common.TreeNode) bool {
	return height(root) != -1
}

func height(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := height(root.Right)
	if rightHeight == -1 || abs(leftHeight, rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1

}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}

	return a - b
}
