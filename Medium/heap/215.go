package heap

import (
	"fmt"
)

func FindKthLargestTest() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 1},
		{[]int{99, 99}, 1, 99},
	}

	for idx, tc := range testCases {
		result := findKthLargest(tc.nums, tc.k)
		if result == tc.expected {
			fmt.Printf("Test Case %d: Passed - Got %d\n", idx, result)
		} else {
			fmt.Printf("Test Case %d: Failed - Got %d, Expected %d\n", idx, result, tc.expected)
		}
	}
}

func findKthLargest(nums []int, k int) int {
	return 0
}
