package Array

import (
	"testing"
)

func TestFindRepeatDocument(t *testing.T) {
	testCases := []struct {
		name      string
		documents []int
		expected  int
	}{
		{
			name:      "first repeat",
			documents: []int{1, 2, 3, 4, 2, 5},
			expected:  2,
		},
		{
			name:      "no repeats",
			documents: []int{1, 2, 3, 4, 5},
			expected:  -1, // assuming -1 as an indicator of no repeats
		},
		{
			name:      "multiple repeats",
			documents: []int{1, 2, 2, 3, 3},
			expected:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findRepeatDocument(tc.documents)
			if actual != tc.expected {
				t.Errorf("Failed '%s': expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func findRepeatDocument(documents []int) int {
	docMap := make(map[int]int)

	for _, document := range documents {
		if docMap[document] == 1 {
			return document
		}
		docMap[document]++
	}

	return -1
}
