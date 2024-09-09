package Array

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Test1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Test2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Test3",
			nums: []int{0},
			want: true,
		},
		{
			name: "Test4",
			nums: []int{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump2(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func canJump2(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	maxJump := nums[0]
	for i := 0; i < maxJump; i++ {
		if i+nums[i] >= len(nums)-1 {
			return true
		}

		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
	}

	return false
}
