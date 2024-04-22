package dp

import (
	"fmt"
)

func JumpTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 2},
	}

	for _, tc := range testCases {
		result := jump1(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Got %v\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %v, Got %v\n", tc.nums, tc.expected, result)
		}
	}
}

// dp
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range len(nums) - 1 {
		dp[i] = i
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

// greedy
func jump1(nums []int) int {
	start := 0
	end := 1
	counts := 0
	for end < len(nums) {
		maxVal := 0
		for start < end {
			maxVal = max(start+nums[start], maxVal)
			start++
		}
		start = end
		end = maxVal
		counts++
	}

	return counts
}
