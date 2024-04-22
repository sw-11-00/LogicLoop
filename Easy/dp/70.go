package dp

import (
	"fmt"
)

func ClimbStairsTest() {
	testCases := []struct {
		n        int // 输入n阶楼梯
		expected int // 预期的方法数量
	}{
		{1, 1},   // 只有1种方式爬1阶
		{2, 2},   // 两种方式爬2阶：1+1和2
		{3, 3},   // 三种方式爬3阶：1+1+1, 1+2, 2+1
		{4, 5},   // 五种方式爬4阶：1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2
		{5, 8},   // 八种方式爬5阶
		{10, 89}, // 89种方式爬10阶
	}

	for _, tc := range testCases {
		result := climbStairs(tc.n)
		if result == tc.expected {
			fmt.Printf("Test Passed for %d stairs: Expected and got %d\n", tc.n, result)
		} else {
			fmt.Printf("Test Failed for %d stairs: Expected %d, Got %d\n", tc.n, tc.expected, result)
		}
	}
}

func climbStairs(n int) int {
	s := make([]int, n+1)
	return climbStairsStore(n, s)
}

func climbStairsStore(n int, memo []int) int {
	if memo[n] > 0 {
		return memo[n]
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	memo[n] = climbStairsStore(n-1, memo) + climbStairsStore(n-2, memo)
	return memo[n]
}
