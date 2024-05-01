package strings

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Test1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Test2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Test3",
			s:    "",
			want: true,
		},
		{
			name: "Test4",
			s:    "a",
			want: true,
		},
		{
			name: "Test5",
			s:    "ab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func valid(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
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
