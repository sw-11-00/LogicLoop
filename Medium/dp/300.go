package dp

import (
	"fmt"
)

func LengthOfLISTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // 最长递增子序列为[2, 3, 7, 101]
		{[]int{0, 1, 0, 3, 2, 3}, 4},           // 最长递增子序列为[0, 1, 2, 3]
		{[]int{7, 7, 7, 7, 7}, 1},              // 最长递增子序列为[7]
		{[]int{}, 0},                           // 空数组，无递增子序列
		{[]int{2, 2, 2, 2, 2}, 1},              // 所有元素相同，递增子序列为任一单元素
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6}, // 最长递增子序列为[1, 3, 6, 7, 9, 10]
	}

	for _, tc := range testCases {
		result := lengthOfLIS(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Expected and got %d\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %d, Got %d\n", tc.nums, tc.expected, result)
		}
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}
