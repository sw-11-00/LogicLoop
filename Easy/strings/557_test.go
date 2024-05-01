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
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reverseWords(s string) string {
	sByte := []byte(s)
	i := 0
	var ansByte []byte
	for j := 0; j <= len(sByte); {
		if j < len(sByte) && sByte[j] != ' ' {
			j++
			continue
		}

		ans := reverse(sByte[i:j])
		ansByte = append(ansByte, ans...)
		ansByte = append(ansByte, ' ')
		i = j + 1
		j = j + 1
	}

	return string(ansByte[:len(ansByte)-1])
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
