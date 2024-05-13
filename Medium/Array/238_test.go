package Array

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Test2",
			nums: []int{0, 0},
			want: []int{0, 0},
		},
		{
			name: "Test3",
			nums: []int{1, 0},
			want: []int{0, 1},
		},
		{
			name: "Test4",
			nums: []int{-1, 1, 0},
			want: []int{0, 0, -1},
		},
		{
			name: "Test5",
			nums: []int{-1, -1, -1, -1},
			want: []int{-1, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}

	leftSum := 1
	rightSum := 1
	for i, j := 0, len(nums)-1; i <= len(nums)-1; {
		ans[i] *= leftSum
		ans[j] *= rightSum

		leftSum *= nums[i]
		rightSum *= nums[j]

		i++
		j--
	}

	return ans
}
