package dp

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Test1",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "Test2",
			n:    1,
			want: []string{"()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				return
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {
		cur := make([]string, 0)
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					cur = append(cur, "("+left+")"+right)
				}
			}
		}
		dp[i] = cur
	}

	return dp[n]
}
