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
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "Test2",
			s:    "God is good",
			want: "doG si doog",
		},
		{
			name: "Test3",
			s:    "",
			want: "",
		},
		{
			name: "Test4",
			s:    "Hello",
			want: "olleH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords1(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}

func reverseWords1(s string) string {
	sByte := []byte(s)
	res := ""
	j := 0
	i := 0
	for i <= len(sByte) {
		if i < len(sByte) && sByte[i] != ' ' {
			i++
			continue
		}

		res += string(reverse(sByte[j:i]))
		res += " "
		j = i + 1
		i = i + 1
	}

	return res[:len(res)-1]
}
