package strings

import (
	"testing"
)

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"Test 1", "00110011", 6},
		{"Test 2", "10101", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countBinarySubstrings(tt.s)
			if result != tt.expected {
				t.Errorf("countBinarySubstrings() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func countBinarySubstrings(s string) int {
	count := 0
	prev := 0
	curr := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			count += min(prev, curr)
			prev = curr
			curr = 1
		}
	}

	return count + min(prev, curr)
}
