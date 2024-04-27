package bit

import (
	"math/big"
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
			if got := addBinary1(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary(%s, %s) = %s, want %s", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func addBinary(a string, b string) string {
	aBigInt, success := new(big.Int).SetString(a, 2)
	if !success {

	}
	bBigInt, success := new(big.Int).SetString(b, 2)
	if !success {

	}
	return new(big.Int).Add(aBigInt, bBigInt).Text(2)
}

func addBinary1(a string, b string) string {
	ans := ""
	add := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(a[i]) - '0'
		}
		if j >= 0 {
			y = int(b[j]) - '0'
		}
		add += x + y
		ans = strconv.Itoa(add&1) + ans
		add >>= 1
	}

	return ans
}
