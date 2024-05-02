package Array

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Test1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Test2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Test3",
			strs: []string{"ab", "a"},
			want: "a",
		},
		{
			name: "Test4",
			strs: []string{"a"},
			want: "a",
		},
		{
			name: "Test5",
			strs: []string{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	commonStr := comp(strs[0], strs[1])
	j := 2
	for j < len(strs) {
		commonStr = comp(commonStr, strs[j])
		if commonStr == "" {
			return commonStr
		}
		j++
	}

	return commonStr
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func comp(str0, str1 string) string {
	commonStr := ""
	minLen := min(len(str0), len(str1))
	for i := 0; i < minLen; i++ {
		if str0[i] == str1[i] {
			commonStr += string(str0[i])
		} else {
			return commonStr
		}
	}

	return commonStr
}
