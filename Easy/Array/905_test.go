package Array

import (
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test1",
			nums: []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
		{
			name: "test2",
			nums: []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.nums); !IsIntArrayEqual(got, tt.want) {
				t.Errorf("sortArrayByParity(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func IsIntArrayEqual(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}

func sortArrayByParity(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]%2 == 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}

	return nums
}
