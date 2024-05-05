package tree

import (
	"reflect"
	"testing"

	"LogicLoop/common"
)

func TestBuildTree1(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *common.TreeNode
	}{
		{
			name:     "empty",
			preorder: []int{},
			inorder:  []int{},
			want:     nil,
		},
		{
			name:     "single element",
			preorder: []int{1},
			inorder:  []int{1},
			want:     &common.TreeNode{Val: 1},
		},
		{
			name:     "multiple elements",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
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
			if got := buildTree1(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree1(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var rootIndex int
	root := &common.TreeNode{Val: preorder[0]}
	for idx, val := range inorder {
		if val == preorder[0] {
			rootIndex = idx
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}
