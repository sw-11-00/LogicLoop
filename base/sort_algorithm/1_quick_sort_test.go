package sort_algorithm

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"demo", []int{21, 25, 49, 26, 16, 8}, []int{8, 16, 21, 25, 26, 49}},
		{"empty slice", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"two elements", []int{2, 1}, []int{1, 2}},
		{"multiple elements", []int{4, 2, 5, 3, 1}, []int{1, 2, 3, 4, 5}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := QuickSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Failed %s, expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	QuickSort1(nums, 0, len(nums)-1)
	return nums
}

func QuickSort1(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		QuickSort1(nums, 0, pivot-1)
		QuickSort1(nums, pivot+1, end)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] < pivot {
			low++
		}
		nums[high] = nums[low]
	}

	nums[low] = pivot
	return low
}
