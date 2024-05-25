package tree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *common.TreeNode
		expected [][]int
	}{
		{"empty tree", nil, [][]int{}},
		{"single node", &common.TreeNode{Val: 1}, [][]int{{1}}},
		{"multiple nodes", &common.TreeNode{
			Val:  3,
			Left: &common.TreeNode{Val: 9},
			Right: &common.TreeNode{
				Val:   20,
				Left:  &common.TreeNode{Val: 15},
				Right: &common.TreeNode{Val: 7},
			},
		}, [][]int{{3}, {9, 20}, {15, 7}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*common.TreeNode{root}
	var result [][]int
	for len(queue) > 0 {
		var level []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
