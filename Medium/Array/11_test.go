package Array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Simple case",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Decreasing height",
			height:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			expected: 8,
		},
		{
			name:     "Increasing height",
			height:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: 8,
		},
		{
			name:     "Uniform height",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxArea(tc.height)
			if actual != tc.expected {
				t.Errorf("Test '%s' failed: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
func maxArea(height []int) int {
	return 0
}
