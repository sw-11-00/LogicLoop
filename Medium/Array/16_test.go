package Array

import (
	"sort"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 0, 0},
		{"single element", []int{1}, 1, 1},
		{"multiple elements", []int{-1, 2, 1, -4}, 1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSumClosest(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			currentSum := nums[i] + nums[j] + nums[k]
			if currentSum == target {
				return target
			}
			if abs(currentSum-target) < abs(sum-target) {
				sum = currentSum
			}
			if currentSum < target {
				j++
			} else {
				k--
			}
		}
	}
	return sum
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}

	return i
}
