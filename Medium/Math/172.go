package Math

import (
	"fmt"
)

func TrailingZeroesTest() {
	// 测试用例 0: 输入 5
	testInput0 := 5
	result0 := trailingZeroes(testInput0)
	fmt.Printf("Test Case 0: Input %d -> Output %d\n", testInput0, result0)

	// 测试用例 1: 输入 10
	testInput1 := 10
	result1 := trailingZeroes(testInput1)
	fmt.Printf("Test Case 1: Input %d -> Output %d\n", testInput1, result1)

	// 测试用例 2: 输入 100
	testInput2 := 100
	result2 := trailingZeroes(testInput2)
	fmt.Printf("Test Case 2: Input %d -> Output %d\n", testInput2, result2)

	// 测试用例 3: 输入 0
	testInput3 := 0
	result3 := trailingZeroes(testInput3)
	fmt.Printf("Test Case 3: Input %d -> Output %d\n", testInput3, result3)
}

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	n /= 5
	return n + trailingZeroes(n)
}
