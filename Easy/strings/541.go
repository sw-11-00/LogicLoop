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
	return ""
}
