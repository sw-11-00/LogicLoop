package Array

import (
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "test2",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "test3",
			nums: []int{2, 1, -1},
			want: 0,
		},
		{
			name: "test4",
			nums: []int{0, 0, 0, 0, 1},
			want: 4,
		},
		{
			name: "test5",
			nums: []int{1, 1, 1, 1, 1, 1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("pivotIndex(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	leftSum := 0
	rightSum := 0
	for _, v := range nums {
		rightSum += v
	}

	for i, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}

	return -1
}
