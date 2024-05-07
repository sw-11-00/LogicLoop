package dp

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "Test2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "Test3",
			s:    "a",
			want: "a",
		},
		{
			name: "Test4",
			s:    "ac",
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	start, maxLength := 0, 1
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLength = length
			}
		}
	}

	return s[start : start+maxLength]
}
