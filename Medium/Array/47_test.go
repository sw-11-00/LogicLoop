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
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	var result [][]int
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] == true {
			continue
		}
		visited[nums[i]] = true
		remaing := make([]int, 0)
		remaing = append(remaing, nums[:i]...)
		remaing = append(remaing, nums[i+1:]...)
		perms := permuteUnique(remaing)
		for _, perm := range perms {
			result = append(result, append([]int{nums[i]}, perm...))
		}
	}

	return result
}
