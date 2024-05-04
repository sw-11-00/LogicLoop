package Array

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Positive numbers",
			nums:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "Mixed numbers",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6, // subarray [4,-1,2,1]
		},
		{
			name:     "Single negative number",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-2, -3, -4, -1},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxSubArray(tc.nums)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
