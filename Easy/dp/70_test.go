package dp

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test1",
			n:    1,
			want: 1,
		},
		{
			name: "Test2",
			n:    2,
			want: 2,
		},
		{
			name: "Test3",
			n:    3,
			want: 3,
		},
		{
			name: "Test4",
			n:    4,
			want: 5,
		},
		{
			name: "Test5",
			n:    45,
			want: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
