package strings

import (
	"fmt"
)

func ReverseStringTest() {
	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
		{[]byte("a"), []byte("a")},
		{[]byte(""), []byte("")},
		{[]byte("😊🚀🌟"), []byte("🌟🚀😊")},
	}

	for _, tc := range testCases {
		original := make([]byte, len(tc.input))
		copy(original, tc.input) // Copy input to preserve for output
		reverseString(tc.input)
		if string(tc.input) == string(tc.expected) {
			fmt.Printf("Test Passed for '%s': Got '%s'\n", string(original), string(tc.input))
		} else {
			fmt.Printf("Test Failed for '%s': Expected '%s', Got '%s'\n", string(original), string(tc.expected), string(tc.input))
		}
	}
}

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
