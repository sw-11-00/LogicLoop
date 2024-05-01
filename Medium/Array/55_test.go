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
			if got := canJump1(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 超烂解法
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	dp := map[int]bool{}
	dp[0] = true
	start := 0
	for i := 1; i < len(nums); i++ {
		for j := start; j < i; j++ {
			if dp[j] == true && j+nums[j] >= i {
				dp[i] = true
				start = j
				break
			}
		}
	}

	return dp[len(nums)-1]
}

func canJump1(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	maxJump := nums[0]
	for i := 1; i <= maxJump; i++ {
		if nums[i]+i > len(nums)-1 {
			return true
		}

		if nums[i]+i > maxJump {
			maxJump = nums[i] + i
		}
	}

	return false
}
