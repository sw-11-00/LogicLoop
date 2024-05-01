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
	first := -1 << 63
	second := -1 << 63
	third := -1 << 63

	for _, num := range nums {
		if num == first || num == second || num == third {
			continue
		}
		if num > first {
			third = second
			second = first
			first = num
		} else if num > second {
			third = second
			second = num
		} else if num > third {
			third = num
		}
	}

	if third == -1<<63 {
		return first
	}

	return third
}
