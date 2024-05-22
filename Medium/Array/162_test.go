package Array

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(int) bool
	}{
		{
			name: "single element",
			nums: []int{1},
			validate: func(index int) bool {
				return index == 0
			},
		},
		{
			name: "multiple elements",
			nums: []int{1, 2, 3, 1},
			validate: func(index int) bool {
				return index == 2
			},
		},
		{
			name: "multiple peaks",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			validate: func(index int) bool {
				return index == 1 || index == 5
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement1(tt.nums)
			if !tt.validate(result) {
				t.Errorf("findPeakElement() = %v, want a valid peak index", result)
			}
		})
	}
}

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func findPeakElement1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}

	return len(nums) - 1
}
