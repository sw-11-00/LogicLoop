package strings

import (
	"fmt"
)

func ReverseWordsTest() {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"hello world", "olleh dlrow"},
		{"  hello world  ", "  olleh dlrow  "},
		{"a b c", "a b c"},
		{"", ""},
		{" singleword ", " drowelgnis "},
	}

	for _, tc := range testCases {
		result := reverseWords(tc.input)
		if result == tc.expected {
			fmt.Printf("Test Passed for \"%s\": Got \"%s\"\n", tc.input, result)
		} else {
			fmt.Printf("Test Failed for \"%s\": Expected \"%s\", Got \"%s\"\n", tc.input, tc.expected, result)
		}
	}
}

func reverseWords(s string) string {
	a := []byte(s)
	for i, j := 0, 0; i < len(a); j = i {
		for j < len(a) && a[j] != ' ' {
			j = j + 1
		}
		reverse(a[i:j])
		i = j + 1
	}

	return string(a)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
