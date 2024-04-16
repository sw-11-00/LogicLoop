package strings

import (
	"fmt"
)

func IsSubsequenceTest() {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		//{"abc", "ahbgdc", true},
		//{"axc", "ahbgdc", false},
		//{"", "ahbgdc", true}, // 空字符串是任何字符串的子序列
		//{"abc", "", false},   // 非空字符串不可能是空字符串的子序列
		{"b", "abc", true},
		{"db", "abc", false},
		{"hello", "hello world", true},
		{"hello", "hlelo world", true},
		{"hello", "helol world", false},
	}

	// 循环执行每个测试用例
	for i, tc := range testCases {
		result := isSubsequence(tc.s, tc.t)
		fmt.Printf("Test Case %d: Input S=\"%s\", T=\"%s\" -> Output %t (Expected %t)\n", i, tc.s, tc.t, result, tc.expected)
	}
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	i := 0
	j := 0
	for j < len(t) {
		if s[i] == t[j] {
			if i == len(s)-1 {
				return true
			}
			i++
		}
		j++
	}

	return false
}
