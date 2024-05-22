package Array

import (
	"testing"
)

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{1, 3, 5, 4, 7},
			want: 3,
		},
		{
			name: "test2",
			nums: []int{2, 2, 2, 2, 2},
			want: 1,
		},
		{
			name: "test3",
			nums: []int{1, 3, 5, 7},
			want: 4,
		},
		{
			name: "test4",
			nums: []int{1, 3, 5, 4, 2, 3, 4, 5},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentLen := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLen++
		} else {
			currentLen = 1
		}

		maxLen = max(currentLen, maxLen)
	}

	return maxLen
}
