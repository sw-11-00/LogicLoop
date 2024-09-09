package Array

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{"empty array", []int{}, [][]int{}},
		{"single element", []int{1}, [][]int{{1}}},
		{"multiple elements", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permute(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("permute() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		remaining := make([]int, 0)
		remaining = append(remaining, nums[:i]...)
		remaining = append(remaining, nums[i+1:]...)
		perms := permute(remaining)
		for _, perm := range perms {
			res = append(res, append([]int{nums[i]}, perm...))
		}
	}

	return res

}
