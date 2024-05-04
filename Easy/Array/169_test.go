package Array

import (
	"sort"
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
	cd := nums[0]
	cnt := 1
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i] == cd {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				cd = nums[i]
				cnt = 1
			}
		}
	}
	return cd
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
