package strings

import (
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "example 1",
			s:    "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "example 2",
			s:    "a",
			want: [][]string{
				{"a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.s); !isSame2DStringArray(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isSame2DStringArray(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !isSameStringArray(a[i], b[i]) {
			return false
		}
	}

	return true
}

func isSameStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func partition(s string) [][]string {
	var res [][]string
	var dfs func(s string, start int, path []string)
	dfs = func(s string, start int, path []string) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				path = append(path, s[start:i+1])
				dfs(s, i+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(s, 0, []string{})
	return res
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
