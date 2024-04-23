package BinaryTree

import (
	"testing"

	"LogicLoop/common"
)

func TestIsSymmetric(t *testing.T) {
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
			name: "symmetric tree",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val:   2,
					Left:  &common.TreeNode{Val: 3},
					Right: &common.TreeNode{Val: 4},
				},
				Right: &common.TreeNode{
					Val:   2,
					Left:  &common.TreeNode{Val: 4},
					Right: &common.TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "asymmetric tree",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
				},
				Right: &common.TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSymmetric(tc.root)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %t, got %t", tc.name, tc.expected, actual)
			}
		})
	}
}

func isSymmetric(root *common.TreeNode) bool {
	return check(root, root)
}

func check(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right)
}

func isSymmetric1(root *common.TreeNode) bool {
	root1, root2 := root, root
	var nodeQueue []*common.TreeNode
	nodeQueue = append(nodeQueue, root1)
	nodeQueue = append(nodeQueue, root2)
	for len(nodeQueue) > 0 {
		a, b := nodeQueue[0], nodeQueue[1]
		nodeQueue = nodeQueue[2:]
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		nodeQueue = append(nodeQueue, a.Left)
		nodeQueue = append(nodeQueue, b.Right)

		nodeQueue = append(nodeQueue, a.Right)
		nodeQueue = append(nodeQueue, b.Left)
	}

	return true
}
