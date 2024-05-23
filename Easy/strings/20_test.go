package strings

import (
	"testing"
)

func TestIsValid1(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Test 1", "()", true},
		{"Test 2", "()[]{}", true},
		{"Test 3", "(]", false},
		{"Test 4", "([)]", false},
		{"Test 5", "{[]}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid2(tt.s)
			if result != tt.expected {
				t.Errorf("isValid1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func isValid1(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValid2(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
