package dp

import (
	"fmt"
)

func JumpTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1, 3, 2, 1, 0, 4}, 3},
	}

	for _, tc := range testCases {
		result := jump(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Got %v\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %v, Got %v\n", tc.nums, tc.expected, result)
		}
	}
}

func jump(nums []int) int {
	return 0
}
