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
			got := myPow1(tt.x, tt.n)
			if !closeEnough(got, tt.want) {
				t.Errorf("myPow(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
			}
		})
	}
}

func myPow1(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}

	return half * half * x
}

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
