package Array

import (
	"testing"
)

func TestThirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "Test2",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "Test3",
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			name: "Test4",
			nums: []int{1, 2, 2, 5, 3, 5},
			want: 2,
		},
		{
			name: "Test5",
			nums: []int{1, 1, 2, 2, 3, 3},
			want: 1,
		},
		{
			name: "Test6",
			nums: []int{1, 2, -2147483648},
			want: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
func thirdMax(nums []int) int {
	max1, max2, max3 := -1<<31-1, -1<<31-1, -1<<31-1
	for _, num := range nums {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 && num < max1 {
			max3 = max2
			max2 = num
		} else if num > max3 && num < max2 {
			max3 = num
		}
	}

	if max3 == -1<<31-1 {
		return max1
	}

	return max3
}
