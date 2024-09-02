package Array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "Test2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		{
			name: "Test3",
			nums: []int{},
			want: 0,
		},
		{
			name: "Test4",
			nums: []int{1, 1, 1, 1, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func removeDuplicates1(nums []int) int {
	k := 0
	for i, num := range nums {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		}
		nums[k] = num
		k++
	}

	return k
}

func removeDuplicates(nums []int) int {
	k := 0
	for i, num := range nums {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		}
		nums[k] = num
	}

	return k
}
