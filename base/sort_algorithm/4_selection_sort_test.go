package sort_algorithm

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"two elements", []int{2, 1}, []int{1, 2}},
		{"multiple elements", []int{4, 2, 5, 3, 1}, []int{1, 2, 3, 4, 5}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SelectionSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Failed %s, expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func SelectionSort(nums []int) []int {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			nums[minIdx], nums[i] = nums[i], nums[minIdx]
		}
	}

	return nums
}
