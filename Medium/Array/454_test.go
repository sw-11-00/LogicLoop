package Array

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		nums3    []int
		nums4    []int
		expected int
	}{
		{"empty arrays", []int{}, []int{}, []int{}, []int{}, 0},
		{"one empty array", []int{1}, []int{}, []int{1}, []int{-2}, 0},
		{"single elements", []int{1}, []int{-1}, []int{1}, []int{-1}, 1},
		{"multiple elements", []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fourSumCount(tt.nums1, tt.nums2, tt.nums3, tt.nums4)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	res := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			res[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			count += res[-(nums3[i] + nums4[j])]
		}
	}

	return count
}
