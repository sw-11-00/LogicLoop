package strings

import (
	"fmt"
)

func ReverseStrTest() {
	testCases := []struct {
		s        string
		k        int
		expected string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
		{"abcdefg", 8, "gfedcba"},
		{"a", 3, "a"},
		{"abcdefg", 1, "abcdefg"},
		{"abcd", 4, "dcba"},
	}

	for _, tc := range testCases {
		result := reverseStr(tc.s, tc.k)
		if result == tc.expected {
			fmt.Printf("Test Passed for '%s' with k=%d: Got '%s'\n", tc.s, tc.k, result)
		} else {
			fmt.Printf("Test Failed for '%s' with k=%d: Expected '%s', Got '%s'\n", tc.s, tc.k, tc.expected, result)
		}
	}
}

func reverseStr(s string, k int) string {
	a := []byte(s)
	for i := 0; i < len(a); i += 2 * k {
		sub := a[i:min(i+k, len(a))]
		for j := 0; j < len(sub)/2; j++ {
			sub[j], sub[len(sub)-j-1] = sub[len(sub)-j-1], sub[j]
		}
	}

	return string(a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
