package sort_algorithm

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
			result := BubbleSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Failed %s, expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func BubbleSort(nums []int) []int {
	n := len(nums)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
		n--
	}

	return nums
}
