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
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = len(nums)
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}
