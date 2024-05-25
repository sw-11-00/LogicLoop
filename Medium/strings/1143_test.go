package strings

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "both strings are empty",
			text1: "",
			text2: "",
			want:  0,
		},
		{
			name:  "one string is empty",
			text1: "",
			text2: "abc",
			want:  0,
		},
		{
			name:  "both strings are the same",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			name:  "both strings are different",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "both strings have common subsequence",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
