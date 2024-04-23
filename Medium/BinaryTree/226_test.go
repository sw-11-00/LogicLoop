package BinaryTree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *common.TreeNode
		expected *common.TreeNode
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name:     "single node",
			root:     &common.TreeNode{Val: 1},
			expected: &common.TreeNode{Val: 1},
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
			expected: &common.TreeNode{
				Val: 1,
				Right: &common.TreeNode{
					Val:   2,
					Left:  &common.TreeNode{Val: 5},
					Right: &common.TreeNode{Val: 4},
				},
				Left: &common.TreeNode{
					Val:   3,
					Right: &common.TreeNode{Val: 6},
					Left:  &common.TreeNode{Val: 7},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := invertTree(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
