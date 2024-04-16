package Array

import (
	"fmt"
)

func BestTimeToBuyAndSellStockTest() {
	// test case 0
	nums := []int{7, 1, 5, 3, 6, 4}
	numsLen := maxProfit(nums)
	fmt.Println(numsLen)
}

func maxProfit(prices []int) int {
	max := 0
	slow := 0
	fast := 1
	if len(prices) == 1 {
		return 0
	}

	for fast < len(prices) {
		if prices[slow] > prices[fast] {
			slow = fast
		} else {
			max = Max(prices[fast]-prices[slow], max)
		}
		fast++
	}

	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
