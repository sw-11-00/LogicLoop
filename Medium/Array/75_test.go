package Array

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			name: "test3",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "test4",
			nums: []int{1, 2, 0},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors1(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("sortColors() got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func sortColors1(nums []int) {
	start := 0
	mid := 0
	end := len(nums) - 1
	for mid <= end {
		if nums[mid] == 0 {
			nums[start], nums[mid] = nums[mid], nums[start]
			mid++
			start++
		} else if nums[mid] == 1 {
			mid++
		} else if nums[mid] == 2 {
			nums[mid], nums[end] = nums[end], nums[mid]
			end--
		}
	}
}
