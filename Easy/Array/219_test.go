package Array

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 4, 5, 6}, 2, false},
		{[]int{}, 0, false}, // 测试空数组
	}

	for _, test := range tests {
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.expected {
			t.Errorf("ContainsNearbyDuplicate(%v, %d) = %v; expected %v", test.nums, test.k, result, test.expected)
		}
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i, num := range nums {
		if preIdx, ok := numMap[num]; ok && i-preIdx <= k {
			return true
		}
		numMap[num] = i
	}

	return false
}
