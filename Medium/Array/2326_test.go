package Array

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	left, right, top, bottom := 0, n-1, 0, m-1
	current := head
	for current != nil {
		for i := left; i <= right && current != nil; i++ {
			res[top][i] = current.Val
			current = current.Next
		}
		top++
		for i := top; i <= bottom && current != nil; i++ {
			res[i][right] = current.Val
			current = current.Next
		}
		right--
		for i := right; i >= left && current != nil; i-- {
			res[bottom][i] = current.Val
			current = current.Next
		}
		bottom--
		for i := bottom; i >= top && current != nil; i-- {
			res[i][left] = current.Val
			current = current.Next
		}
		left++
	}

	return res
}
