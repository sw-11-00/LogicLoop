package Array

import (
	"sort"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},                   // 2 是缺失的数字
		{[]int{0, 1}, 2},                      // 2 是缺失的数字
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8}, // 8 是缺失的数字
		{[]int{0}, 1},                         // 1 是缺失的数字
		{[]int{1}, 0},                         // 0 是缺失的数字
		{[]int{}, 0},                          // 边界情况：空数组时，缺失的数字是 0
	}

	for _, test := range tests {
		result := missingNumber1(test.nums)
		if result != test.expected {
			t.Errorf("MissingNumber(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func missingNumber(nums []int) int {
	sort.Ints(nums)

	start := 0
	for _, num := range nums {
		if num != start {
			return start
		}
		start++
	}

	return start
}

func missingNumber1(nums []int) int {
	n := len(nums)
	numSum := n * (n + 1) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return numSum - sum
}
