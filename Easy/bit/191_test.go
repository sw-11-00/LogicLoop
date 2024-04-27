package bit

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Binary 101010", 0b101010, 3},
		{"MaxInt", 0x7FFFFFFF, 31},   // 2147483647 in binary has 31 ones
		{"FullBits", 0xFFFFFFFF, 32}, // -1 in a two's complement 32-bit system, all bits are 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.n); got != tt.want {
				t.Errorf("%s failed: got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

func hammingWeight(n int) int {
	weight := 0
	for i := 0; i < 32; i++ {
		if 1<<i&n > 0 {
			weight++
		}
	}

	return weight
}
