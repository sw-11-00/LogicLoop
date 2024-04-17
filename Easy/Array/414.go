package Array

import (
	"fmt"
)

func ThirdMaxTest() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1, 2, 2, 3}, 1},
		{[]int{1, 2, 2, 5, 3, 5}, 2},
		{[]int{5, 2, 2}, 5},
		{[]int{1, 2, -2147483648}, -2147483648},
	}

	for idx, tc := range testCases {
		result := thirdMax(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test Case %d: Passed - Got %d\n", idx, result)
		} else {
			fmt.Printf("Test Case %d: Failed - Got %d, Expected %d\n", idx, result, tc.expected)
		}
	}
}

func thirdMax(nums []int) int {
	first := -9223372036854775808
	second := -9223372036854775808
	third := -9223372036854775808

	for _, num := range nums {
		if num > first {
			third = second
			second = first
			first = num
		} else if num > second && num < first {
			third = second
			second = num
		} else if num > third && num < second {
			third = num
		}
	}

	if first == -9223372036854775808 || second == -9223372036854775808 || third == -9223372036854775808 {
		return first
	}

	return third
}
