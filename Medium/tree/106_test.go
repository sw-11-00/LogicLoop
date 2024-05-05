package tree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      *common.TreeNode
	}{
		{
			name:      "empty",
			inorder:   []int{},
			postorder: []int{},
			want:      nil,
		},
		{
			name:      "single element",
			inorder:   []int{1},
			postorder: []int{1},
			want:      &common.TreeNode{Val: 1},
		},
		{
			name:      "multiple elements",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &common.TreeNode{
				Val: 3,
				Left: &common.TreeNode{
					Val: 9,
				},
				Right: &common.TreeNode{
					Val: 20,
					Left: &common.TreeNode{
						Val: 15,
					},
					Right: &common.TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree(inorder []int, postorder []int) *common.TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// Find the root node
	rootVal := postorder[len(postorder)-1]
	root := &common.TreeNode{Val: rootVal}

	// Find the index of the root node in the inorder list
	var rootIndex int
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	// Recursively build the left and right subtrees
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}
