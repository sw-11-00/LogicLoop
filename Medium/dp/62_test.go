package dp

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "Test1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "Test2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "Test3",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "Test4",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "Test5",
			m:    1,
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
