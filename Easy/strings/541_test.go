package strings

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "Test1",
			s:    "abcdefg",
			k:    2,
			want: "bacdfeg",
		},
		{
			name: "Test2",
			s:    "abcd",
			k:    4,
			want: "dcba",
		},
		{
			name: "Test3",
			s:    "",
			k:    1,
			want: "",
		},
		{
			name: "Test4",
			s:    "a",
			k:    2,
			want: "a",
		},
		{
			name: "Test5",
			s:    "abcde",
			k:    1,
			want: "abcde",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.s, tt.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reverseStr(s string, k int) string {
	sByte := []byte(s)
	for i := 0; i < len(sByte); i += 2 * k {
		left := i
		right := min(i+k-1, len(sByte)-1)
		for left < right {
			sByte[left], sByte[right] = sByte[right], sByte[left]
			left++
			right--
		}
	}
	return string(sByte)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
