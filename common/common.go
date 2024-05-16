package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqualTreeNode(got *TreeNode, expected *TreeNode) bool {
	if got == nil && expected == nil {
		return true
	}
	if got == nil || expected == nil {
		return false
	}
	return got.Val == expected.Val && IsEqualTreeNode(got.Left, expected.Left) && IsEqualTreeNode(got.Right, expected.Right)
}
