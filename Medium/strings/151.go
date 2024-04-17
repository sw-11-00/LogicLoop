package strings

import (
	"fmt"
)

func ReverseWordsTest() {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "olleh dlrow"},
		{"  hello world  ", "  olleh dlrow  "},
		{"a b c", "a b c"},
		{"Example   good", "elpmaxE   doog"},
		{"", ""},
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

}
