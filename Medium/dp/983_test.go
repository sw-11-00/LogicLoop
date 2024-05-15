package dp

import (
	"fmt"
	"testing"
)

func TestMinCostTickets(t *testing.T) {
	testCases := []struct {
		days     []int
		costs    []int
		expected int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},                                          // Example 1: 使用2天单日票，1个七日票
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30}, []int{2, 7, 15}, 17},                          // Example 2: 使用2个七日票，1个单日票
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}, []int{3, 14, 50}, 30}, // Example 3: 使用2个月票
		{[]int{6, 8, 9, 18, 20, 21, 23, 25}, []int{10, 20, 50}, 30},                              // Example 4: 使用3个七日票
		{[]int{1, 100, 200}, []int{1, 10, 100}, 3},                                               // 长时间间隔，每日票更便宜
	}

	for _, tc := range testCases {
		result := minCostTickets(tc.days, tc.costs)
		if result == tc.expected {
			fmt.Printf("Test Passed for days %v with costs %v: Expected and got %d\n", tc.days, tc.costs, result)
		} else {
			fmt.Printf("Test Failed for days %v with costs %v: Expected %d, Got %d\n", tc.days, tc.costs, tc.expected, result)
		}
	}
}

func minCostTickets(days []int, costs []int) int {
	return 0
}
