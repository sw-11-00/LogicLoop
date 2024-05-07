package Array

import (
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			name: "Test2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 9,
		},
		{
			name: "Test3",
			nums: []int{},
			want: 0,
		},
		{
			name: "Test4",
			nums: []int{1, 1, 1, 1, 1, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates2(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func removeDuplicates2(nums []int) int {
	k := 0
	for _, num := range nums {
		if k < 2 || nums[k-2] < num {
			nums[k] = num
			k++
		}
	}

	return k
}
