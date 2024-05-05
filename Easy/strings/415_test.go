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
		{"Different length", "123", "9876", "9999"},
		{"Large numbers", "123456789", "987654321", "1111111110"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.num1, tt.num2); got != tt.want {
				t.Errorf("addStrings(%s, %s) = %s, want %s", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 && len(num2) == 0 {
		return "0"
	}
	ans := ""
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i]) - '0'
		}
		if j >= 0 {
			y = int(num2[j]) - '0'
		}
		sum := x + y + carry
		carry = sum / 10
		ans = strconv.Itoa(sum%10) + ans
		i--
		j--
	}

	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}

	return ans
}
