package dp

import (
	"fmt"
)

func FindNumberOfLISTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5, 4, 7}, 2},              // 最长递增子序列有两个：[1, 3, 5, 7] 和 [1, 3, 4, 7]
		{[]int{2, 2, 2, 2, 2}, 5},              // 所有元素相同，每个元素都可以是一个单独的子序列
		{[]int{1, 2, 4, 3, 5, 4, 7, 2}, 3},     // 最长递增子序列有三个
		{[]int{}, 0},                           // 空数组，没有子序列
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 1}, // 最长递增子序列只有一个：[2, 3, 7, 18, 101]
	}

	for _, tc := range testCases {
		result := findNumberOfLIS(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Expected and got %d\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %d, Got %d\n", tc.nums, tc.expected, result)
		}
	}
}

// hard
func findNumberOfLIS(nums []int) int {
	return 0
}
