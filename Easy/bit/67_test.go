package bit

import (
	"strconv"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"Adding mixed", "1010", "1011", "10101"},
		{"Adding different lengths", "1", "111", "1000"},
		{"Adding with carry", "11", "1", "100"},
		{"Adding zeros", "0", "0", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary(%s, %s) = %s, want %s", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func addBinary(a string, b string) string {
	if len(a) == 0 && len(b) == 0 {
		return "0"
	}
	ans := ""
	carry := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i]) - '0'
		}
		if j >= 0 {
			y = int(b[j]) - '0'
		}
		sum := x + y + carry
		carry = sum / 2
		ans = strconv.Itoa(sum%2) + ans
		i--
		j--
	}

	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}

	return ans
}
