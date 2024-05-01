package strings

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Test1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "Test2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "Test3",
			s:    "",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "Test4",
			s:    "abc",
			t:    "",
			want: false,
		},
		{
			name: "Test5",
			s:    "",
			t:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.s, tt.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isSubsequence(s string, t string) bool {
	sByte := []byte(s)
	tByte := []byte(t)

	j := 0
	for i := 0; i < len(tByte); i++ {
		if j < len(sByte) && sByte[j] == tByte[i] {
			j++
		}
	}

	if j == len(s) {
		return true
	}

	return false
}
