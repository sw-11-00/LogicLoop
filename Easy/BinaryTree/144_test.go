package BinaryTree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *common.TreeNode
		expected []int
	}{
		{"empty tree", nil, []int{}},
		{"single node", &common.TreeNode{Val: 1}, []int{1}},
		{"multiple nodes", &common.TreeNode{
			Val:   1,
			Left:  &common.TreeNode{Val: 2},
			Right: &common.TreeNode{Val: 3},
		}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := preorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func preorderTraversal(root *common.TreeNode) []int {
	var result []int
	preorderTraversalHelper(root, &result)
	return result
}

func preorderTraversalHelper(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preorderTraversalHelper(root.Left, result)
	preorderTraversalHelper(root.Right, result)
}
