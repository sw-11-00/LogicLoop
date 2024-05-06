package bit

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	testCases := []struct {
		x      int
		y      int
		expect int
	}{
		{x: 1, y: 4, expect: 2},
		{x: 3, y: 1, expect: 1},   // 3 (011) and 1 (001) have 1 bit different
		{x: 14, y: 15, expect: 1}, // 14 (1110) and 15 (1111) have 1 bit different
		{x: 7, y: 0, expect: 3},   // 7 (0111) and 0 (0000) have 3 bits different
		{x: 0, y: 0, expect: 0},   // 0 (0000) and 0 (0000) have 0 bits different
		{x: 8, y: 15, expect: 3},  // 8 (1000) and 15 (1111) have 3 bits different
	}

	for _, tt := range testCases {
		result := hammingDistance1(tt.x, tt.y)
		if result != tt.expect {
			t.Errorf("hammingDistance(%v, %v) = %v; want %v", tt.x, tt.y, result, tt.expect)
		}
	}
}

func hammingDistance1(x int, y int) int {
	count := 0
	for x > 0 || y > 0 {
		a := x & 1
		b := y & 1
		count += a ^ b
		x >>= 1
		y >>= 1
	}

	return count
}
