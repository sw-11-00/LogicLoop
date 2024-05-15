package strings

import (
	"strconv"
	"testing"
)

func TestAddStrings(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{"Both empty", "", "", "0"},
		{"First empty", "", "123", "123"},
		{"Second empty", "456", "", "456"},
		{"Same length", "123", "876", "999"},
		{"Third length", "1", "9", "10"},
		{"Different length", "123", "9876", "9999"},
		{"Large numbers", "123456789", "987654321", "1111111110"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings1(tt.num1, tt.num2); got != tt.want {
				t.Errorf("addStrings(%s, %s) = %s, want %s", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}

func addStrings1(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	}

	if num2 == "" {
		return num1
	}

	carry := 0
	i := len(num1) - 1
	j := len(num2) - 1
	res := ""
	for i >= 0 || j >= 0 {
		var a int
		var b int
		if i >= 0 {
			a = int(num1[i]) - '0'
		}
		if j >= 0 {
			b = int(num2[j]) - '0'
		}
		sum := a + b + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
		i--
		j--
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}
