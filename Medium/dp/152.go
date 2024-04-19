package dp

import (
	"fmt"
)

func MaxProductTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, 1}, 1},
		{[]int{-2, -3, 7}, 42},
		{[]int{0, 2, 3, -2, 4}, 6},
		{[]int{-1, -2, -3, 0}, 6},
		{[]int{-1, 2, -3, 4, -5}, 120},
		{[]int{3, -1, 4}, 4},
		{[]int{2, -5, -2, -4, 3}, 24},
	}

	for _, tc := range testCases {
		result := maxProduct1(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Expected and got %d\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %d, Got %d\n", tc.nums, tc.expected, result)
		}
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := -111111111111
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum > max {
			max = sum
		}
		j := i + 1
		for j < len(nums) {
			sum *= nums[j]
			if sum > max {
				max = sum
			}
			j++
		}
	}

	return max
}

func maxProduct1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	globalMax := nums[0]
	localMax := nums[0]
	localMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			localMax, localMin = localMin, localMax
		}

		localMax = max(localMax*nums[i], nums[i])
		localMin = min(localMin*nums[i], nums[i])
		globalMax = max(globalMax, localMax)
	}

	return globalMax
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
