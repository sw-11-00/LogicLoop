package LinkedList

import (
	"testing"

	"LogicLoop/common"
)

func TestConvertBiNode(t *testing.T) {
	// root = [4,2,5,1,3]
	// Output: [1,null,2,null,3,null,4,null,5]
	root := &common.TreeNode{Val: 4}
	root.Left = &common.TreeNode{Val: 2}
	root.Right = &common.TreeNode{Val: 5}
	root.Left.Left = &common.TreeNode{Val: 1}
	root.Left.Right = &common.TreeNode{Val: 3}
	expected := &common.TreeNode{Val: 1}
	expected.Right = &common.TreeNode{Val: 2}
	expected.Right.Right = &common.TreeNode{Val: 3}
	expected.Right.Right.Right = &common.TreeNode{Val: 4}
	expected.Right.Right.Right.Right = &common.TreeNode{Val: 5}
	got := convertBiNode(root)
	if !common.IsEqualTreeNode(got, expected) {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

	// root = [2,1,3]
	// Output: [1,null,2,null,3]
	root = &common.TreeNode{Val: 2}
	root.Left = &common.TreeNode{Val: 1}
	root.Right = &common.TreeNode{Val: 3}
	expected = &common.TreeNode{Val: 1}
	expected.Right = &common.TreeNode{Val: 2}
	expected.Right.Right = &common.TreeNode{Val: 3}
	got = convertBiNode(root)
	if !common.IsEqualTreeNode(got, expected) {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func convertBiNode(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	head := &common.TreeNode{}
	cur := head
	stack := []*common.TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Right = node
		cur = cur.Right
		node.Left = nil
		node = node.Right
	}
	return head.Right
}
