package bit

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want uint32
	}{
		{"All zeros", 0, 0},
		{"All ones", 0xFFFFFFFF, 0xFFFFFFFF},
		{"Example 1", 0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{"Example 2", 0b11111111111111111111111111111101, 0b10111111111111111111111111111111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.num)
			if got != tt.want {
				t.Errorf("%s failed: got %08x, want %08x", tt.name, got, tt.want)
			}
		})
	}
}

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}

	return ans
}
