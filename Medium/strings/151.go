package strings

import (
	"fmt"
)

func ReverseWordsTest() {
	testCases := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, tc := range testCases {
		result := reverseWords(tc.input)
		if result == tc.expected {
			fmt.Printf("Test Passed for '%s': Got '%s'\n", tc.input, result)
		} else {
			fmt.Printf("Test Failed for '%s': Expected '%s', Got '%s'\n", tc.input, tc.expected, result)
		}
	}
}

func reverseWords(s string) string {
	res := ""
	a := []byte(s)
	for slow, fast := len(a)-1, len(a)-1; fast > 0; {
		for slow >= 0 && a[slow] == ' ' {
			slow--
		}
		fast = slow
		for fast >= 0 && a[fast] != ' ' {
			fast--
		}
		if fast == slow {
			break
		}
		res = res + string(a[fast+1:slow+1])
		res = res + " "
		slow = fast
	}
	res = res[0 : len(res)-1]
	return res
}
