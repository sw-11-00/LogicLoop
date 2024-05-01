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
			sortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("sortColors() got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
