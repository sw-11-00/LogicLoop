package Array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Target present multiple times",
			nums:     []int{2, 2},
			target:   3,
			expected: []int{-1, -1},
		},
		{
			name:     "Target present multiple times",
			nums:     []int{1, 3},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "Target present multiple times",
			nums:     []int{2, 2},
			target:   2,
			expected: []int{0, 1},
		},
		{
			name:     "Target present multiple times",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "Target present once",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   7,
			expected: []int{1, 2},
		},
		{
			name:     "Target not present",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   1,
			expected: []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := searchRange(tc.nums, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

func searchRange(nums []int, target int) []int {
	result := make([]int, 2)
	result[0], result[1] = -1, -1
	if len(nums) == 0 {
		return result
	}
	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] == target {
			result[0] = mid
			j = mid - 1
		}
	}

	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] == target {
			result[1] = mid
			i = mid + 1
		}
	}

	return result
}
