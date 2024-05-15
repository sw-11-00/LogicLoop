package Math

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for _, test := range tests {
		ret := fib(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
