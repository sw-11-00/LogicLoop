package Array

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{2},
			expected: 2,
		},
		{
			name:     "multiple elements",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2},
			expected: -2,
		},
		{
			name:     "mixed numbers",
			nums:     []int{0, 1, 0, 1, 99},
			expected: 99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := singleNumber(tc.nums)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func singleNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}

	return sum
}
