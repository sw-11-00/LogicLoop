package dp

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Test1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "Test2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "Test3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "Test4",
			coins:  []int{1},
			amount: 1,
			want:   1,
		},
		{
			name:   "Test5",
			coins:  []int{1},
			amount: 2,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] <= amount {
		return dp[amount]
	} else {
		return -1
	}
}
