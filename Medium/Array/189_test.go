package Array

import (
	"reflect"
	"testing"
)

func TestRotateTest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Test1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Test2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			name: "Test3",
			nums: []int{1, 2},
			k:    3,
			want: []int{2, 1},
		},
		{
			name: "Test4",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    7,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Test5",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    0,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Test6",
			nums: []int{},
			k:    1,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
