package Array

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "all zeroes",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "mixed zeroes and numbers",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "no zeroes",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			moveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, tc.nums)
			}
		})
	}
}

func moveZeroes(nums []int) {
	i := 0
	for _, num := range nums {
		if num != 0 {
			nums[i] = num
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
