package strings

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Test 1", "aba", true},
		{"Test 2", "abca", true},
		{"Test 3", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("validPalindrome() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome1(s, l+1, r) || isPalindrome1(s, l, r-1)
		}
		l++
		r--
	}

	return true
}

func isPalindrome1(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
