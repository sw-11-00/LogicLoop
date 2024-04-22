package dp

import (
	"fmt"
)

func CoinChangeTest() {
	testCases := []struct {
		coins    []int // 硬币面额
		amount   int   // 需要找零的金额
		expected int   // 预期的最小硬币数
	}{
		{[]int{1, 2, 5}, 11, 3},              // 5 + 5 + 1 = 11
		{[]int{2}, 3, -1},                    // 无法组成3
		{[]int{1}, 0, 0},                     // 金额为0，需要0个硬币
		{[]int{1, 2, 5}, 100, 20},            // 5*20 = 100
		{[]int{3, 7, 405, 436}, 8839, 25},    // 需要复杂组合
		{[]int{186, 419, 83, 408}, 6249, 20}, // 更大的数
	}

	for _, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)
		if result == tc.expected {
			fmt.Printf("Test Passed for coins %v with amount %d: Expected and got %d\n", tc.coins, tc.amount, result)
		} else {
			fmt.Printf("Test Failed for coins %v with amount %d: Expected %d, Got %d\n", tc.coins, tc.amount, tc.expected, result)
		}
	}
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 << 30
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] < amount {
		return dp[amount]
	} else {
		return -1
	}
}
