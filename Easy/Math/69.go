package Math

import (
	"fmt"
)

func MySqrtTest() {
	// 测试用例 0: 完全平方数
	testInput0 := 4
	result0 := mySqrt(testInput0)
	fmt.Printf("Test Case 0: Input %d -> Output %d\n", testInput0, result0)

	// 测试用例 1: 非完全平方数
	testInput1 := 8
	result1 := mySqrt(testInput1)
	fmt.Printf("Test Case 1: Input %d -> Output %d\n", testInput1, result1)

	// 测试用例 2: 边界情况
	testInput2 := 0
	result2 := mySqrt(testInput2)
	fmt.Printf("Test Case 2: Input %d -> Output %d\n", testInput2, result2)

	testInput3 := 1
	result3 := mySqrt(testInput3)
	fmt.Printf("Test Case 3: Input %d -> Output %d\n", testInput3, result3)

}

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left - 1
}
