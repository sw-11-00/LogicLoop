package BinaryTree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestPostorderTraversal(t *testing.T) {
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
		}, []int{2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := postorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func postorderTraversal(root *common.TreeNode) []int {
	var result []int
	postorderHelper(root, &result)
	return result
}

func postorderHelper(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}

	postorderHelper(root.Left, result)
	postorderHelper(root.Right, result)
	*result = append(*result, root.Val)
}
