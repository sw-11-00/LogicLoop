package Array

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "simple case",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "longer case",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "all same",
			nums:     []int{1, 1, 1, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := majorityElement(tc.nums)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}

	maxElement := numMap[nums[0]]
	ans := nums[0]
	for _, num := range nums {
		if numMap[num] > maxElement {
			maxElement = numMap[num]
			ans = num
		}
	}
	return ans
}
