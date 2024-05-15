package Math

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
	}

	for _, test := range tests {
		if got := isHappy(test.n); got != test.want {
			t.Errorf("isHappy(%v) = %v", test.n, got)
		}
	}
}

func isHappy(n int) bool {
	slow := n
	fast := n
	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(fast)
		fast = sumOfSquares(fast)
		if slow == fast {
			break
		}
	}

	return slow == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
