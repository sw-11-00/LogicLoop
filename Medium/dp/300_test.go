package dp

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "Test2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "Test3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "Test4",
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			want: 6,
		},
		{
			name: "Test5",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}
