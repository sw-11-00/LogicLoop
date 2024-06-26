package strings

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "Test2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "Test3",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "Test4",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Test5",
			s:    "Alice does not even like bob",
			want: "bob like even not does Alice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reverseWords(s string) string {
	sByte := []byte(s)
	fast := len(sByte) - 1
	slow := len(sByte) - 1
	res := ""

	for fast >= 0 {
		for fast >= 0 && sByte[fast] == ' ' {
			fast--
			continue
		}

		if fast == -1 {
			break
		}
		slow = fast

		for fast >= 0 && sByte[fast] != ' ' {
			fast--
		}

		res += string(sByte[fast+1 : slow+1])
		res += " "
		fast--
		slow = fast
	}

	return res[:len(res)-1]
}
