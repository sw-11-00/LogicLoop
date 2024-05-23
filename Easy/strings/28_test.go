package strings

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{"Test 1", "hello", "ll", 2},
		{"Test 2", "aaaaa", "bba", -1},
		{"Test 3", "", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+len(needle) > len(haystack) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}
