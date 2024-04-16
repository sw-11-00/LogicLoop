package Math

import (
	"fmt"
)

func PlusOneTest() {
	// test case 0
	digits := []int{9, 1, 9}
	result := plusOne(digits)
	fmt.Println("Test Case 0: Input [1, 2, 3] -> Output", result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i]%10 != 0 {
			return digits
		} else {
			digits[i] = 0
		}
	}

	return append([]int{1}, digits...)
}
