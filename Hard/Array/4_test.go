package Array

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "both empty",
			nums1:    []int{},
			nums2:    []int{},
			expected: 0.0,
		},
		{
			name:     "first empty",
			nums1:    []int{},
			nums2:    []int{1, 2, 3, 4, 5},
			expected: 3.0,
		},
		{
			name:     "second empty",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{},
			expected: 3.0,
		},
		{
			name:     "both non-empty, same length",
			nums1:    []int{1, 3, 5},
			nums2:    []int{2, 4, 6},
			expected: 3.5,
		},
		{
			name:     "both non-empty, different length",
			nums1:    []int{1, 3},
			nums2:    []int{2, 4, 5, 6},
			expected: 3.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMedianSortedArrays(tc.nums1, tc.nums2)
			if actual != tc.expected {
				t.Errorf("Test %s failed: expected %f, got %f", tc.name, tc.expected, actual)
			}
		})
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0.0
}
