package Array

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Target exists",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Target does not exist - insert position",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Insert at start",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Insert at end",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := searchInsert(tc.nums, tc.target)
			if actual != tc.expected {
				t.Errorf("Failed %s: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if target <= num {
			return i
		}
	}

	return len(nums)
}
