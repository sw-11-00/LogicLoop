package dp

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "Test2",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "Test3",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Test4",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Test5",
			nums: []int{1, 2, 3},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}
