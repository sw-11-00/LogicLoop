package Array

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in, out int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	}

	for _, test := range tests {
		result := reverse1(test.in)
		if result != test.out {
			t.Fatalf("Expected %d but received %d", test.out, result)
		}
	}
}

func reverse1(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
