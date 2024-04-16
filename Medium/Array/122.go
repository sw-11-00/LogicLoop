package Array

import "fmt"

func BestTimeToBuyAndSellStockIITest() {
	// test case 0
	nums := []int{7, 6, 4, 3, 1}
	numsLen := maxProfit(nums)
	fmt.Println(numsLen)
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
