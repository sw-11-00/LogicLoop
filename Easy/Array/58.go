package Array

import (
	"fmt"
)

func LengthOfLastWordTest() {
	testCases := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"", 0},
		{"          ", 0},
		{"single", 6},
		{"a ", 1},
	}

	for _, tc := range testCases {
		result := lengthOfLastWord(tc.s)
		if result == tc.expected {
			fmt.Printf("Test Passed for '%s': Expected and got %d\n", tc.s, result)
		} else {
			fmt.Printf("Test Failed for '%s': Expected %d, Got %d\n", tc.s, tc.expected, result)
		}
	}
}

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	count := 0
	for ; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		} else {
			break
		}
	}

	return count
}
