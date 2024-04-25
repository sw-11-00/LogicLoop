package sort_algorithm

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
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
			result := MergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Failed %s, expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
