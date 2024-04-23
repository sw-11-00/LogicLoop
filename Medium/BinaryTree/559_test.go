package BinaryTree

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func TestNaryMaxDepth(t *testing.T) {
	testCases := []struct {
		name     string
		root     *Node
		expected int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "single node",
			root:     &Node{Val: 1},
			expected: 1,
		},
		{
			name: "three level tree",
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 2,
						Children: []*Node{
							{Val: 5},
							{Val: 6},
						},
					},
					{
						Val: 3,
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 7,
								Children: []*Node{
									{Val: 8},
								},
							},
						},
					},
				},
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NaryMaxDepth(tc.root)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func NaryMaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	childMaxDepth := 0
	for _, child := range root.Children {
		maxDepth := NaryMaxDepth(child)
		if maxDepth > childMaxDepth {
			childMaxDepth = maxDepth
		}
	}

	return childMaxDepth + 1
}
