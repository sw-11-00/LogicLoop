package Array

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Test1",
			s:    "Hello World",
			want: 5,
		},
		{
			name: "Test2",
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "Test3",
			s:    "luffy is still joyboy",
			want: 6,
		},
		{
			name: "Test4",
			s:    "",
			want: 0,
		},
		{
			name: "Test5",
			s:    "a",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count != 0 {
				break
			}
		} else {
			count++
		}
	}
	return count
}
