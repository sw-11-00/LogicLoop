package Array

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{"empty array", []int{}, [][]int{}},
		{"single element", []int{1}, [][]int{{1}}},
		{"multiple elements", []int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permuteUnique(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("permuteUnique() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	var result [][]int
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true

		// Remove the current element from the array
		remaining := make([]int, 0)
		remaining = append(remaining, nums[:i]...)
		remaining = append(remaining, nums[i+1:]...)

		// Recursively call permute on the remaining elements
		perms := permuteUnique(remaining)
		for _, perm := range perms {
			result = append(result, append([]int{nums[i]}, perm...))
		}
	}

	return result
}
