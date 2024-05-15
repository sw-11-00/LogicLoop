package Array

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 2, -1},
	}

	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("search(%v, %v) = %v", test.nums, test.target, got)
		}
	}
}

func search(nums []int, target int) int {
	return 0
}
