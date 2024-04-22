package Array

import (
	"fmt"
)

func LongestConsecutiveTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},              // 最长连续序列是 [1, 2, 3, 4]
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7}, // 最长连续序列是 [-1, 0, 1, 2, 3, 4, 5]
		{[]int{1, 2, 0, 1}, 3},                        // 最长连续序列是 [0, 1, 2]
		{[]int{}, 0},                                  // 空数组
		{[]int{0}, 1},                                 // 单元素数组
		{[]int{9, -8, 2, 6, -1, 0, 1, 5, 8}, 4},       // 最长连续序列是 [-1, 0, 1, 2]
	}

	for _, tc := range testCases {
		result := longestConsecutive(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Passed for %v: Expected and got %d\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %d, Got %d\n", tc.nums, tc.expected, result)
		}
	}
}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	counts := 0
	for i := 0; i < len(nums); i++ {
		if !numSet[nums[i]-1] {
			len := 1
			j := nums[i]
			for numSet[j+1] {
				j++
				len++
			}
			if counts < len {
				counts = len
			}
		}
	}

	return counts
}
