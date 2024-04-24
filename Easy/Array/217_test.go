package Array

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "contains duplicates",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := containsDuplicate(tc.nums)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %t, got %t", tc.name, tc.expected, actual)
			}
		})
	}
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	numMap := make(map[int]int)
	for _, num := range nums {
		if numMap[num] == 1 {
			return true
		}
		numMap[num]++
	}

	return false
}
