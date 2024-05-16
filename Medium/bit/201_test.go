package bit

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{"Range with same numbers", 5, 5, 5},
		{"Range with no common prefix", 5, 7, 4},
		{"Range with different numbers", 0, 1, 0},
		{"Larger range", 10, 15, 8},
		{"Example 1", 26, 30, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rangeBitwiseAnd(tt.left, tt.right)
			if got != tt.want {
				t.Errorf("%s failed: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}

	return left << shift
}

func rangeBitwiseAnd1(left int, right int) {
	shift := 0
	for left != right {
		left >>= 1
		right >>= 1
		shift++
	}

	left <<= shift
}
