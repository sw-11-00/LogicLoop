package Math

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{"Positive power", 2.0, 3, 8.0},
		{"Negative power", 2.0, -3, 0.125},
		{"Zero power", 2.0, 0, 1.0},
		{"Fractional base positive power", 2.5, 2, 6.25},
		{"Fractional base negative power", 2.5, -2, 0.16},
		{"Zero base", 0, 5, 0},
		{"Power of one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myPow(tt.x, tt.n)
			if !closeEnough(got, tt.want) {
				t.Errorf("myPow(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
			}
		})
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	pow := 1.0
	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}

		x *= x
		n >>= 1
	}

	return pow
}

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
