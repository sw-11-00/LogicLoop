package Array

import (
	"fmt"
)

func CanJumpTest() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{2, 5, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
	}

	for _, tc := range testCases {
		result := canJump(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Got %v\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %v, Got %v\n", tc.nums, tc.expected, result)
		}
	}
}

func canJump(nums []int) bool {
	a := len(nums) - 1
	for b := len(nums) - 2; b >= 0; {
		if nums[b]+b >= a {
			a = b
		}
		b--
	}

	if a == 0 {
		return true
	}

	return false
}
