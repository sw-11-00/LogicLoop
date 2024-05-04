package LinkedList

import (
	"testing"

	"LogicLoop/common"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		root *common.TreeNode
	}{
		{
			name: "Test1",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
					Left: &common.TreeNode{
						Val: 3,
					},
					Right: &common.TreeNode{
						Val: 4,
					},
				},
				Right: &common.TreeNode{
					Val: 5,
					Right: &common.TreeNode{
						Val: 6,
					},
				},
			},
		},
		{
			name: "Test2",
			root: &common.TreeNode{
				Val: 1,
				Left: &common.TreeNode{
					Val: 2,
					Left: &common.TreeNode{
						Val: 3,
						Left: &common.TreeNode{
							Val: 7,
						},
						Right: &common.TreeNode{
							Val: 8,
						},
					},
					Right: &common.TreeNode{
						Val: 4,
						Left: &common.TreeNode{
							Val: 9,
						},
						Right: &common.TreeNode{
							Val: 10,
						},
					},
				},
				Right: &common.TreeNode{
					Val: 5,
					Right: &common.TreeNode{
						Val: 6,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
		})
	}
}

func flatten(root *common.TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left != nil {
		right := root.Right
		root.Right = root.Left
		root.Left = nil

		node := root.Right
		for node.Right != nil {
			node = node.Right
		}

		node.Right = right
	}

	return
}
