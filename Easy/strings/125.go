package strings

import (
	"fmt"
	"strings"
)

func IsPalindromeTest() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"Madam, I'm Adam", true},
		{"No lemon, no melon", true},
	}

	// 循环执行每个测试用例
	for i, tc := range testCases {
		result := isPalindrome(tc.input)
		fmt.Printf("Test Case %d: Input \"%s\" -> Output %t (Expected %t)\n", i, tc.input, result, tc.expected)
	}
}

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !valid(s[i]) {
			i++
			continue
		}
		if !valid(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func valid(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}
