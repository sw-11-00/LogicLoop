package Math

import (
	"fmt"
)

func IsPalindromeTest() {
	// test case 0
	nums := 101
	ans := isPalindrome(nums)
	fmt.Println(ans)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	nums := digitsToArray(x)
	fmt.Println(nums)
	start := 0
	end := len(nums) - 1
	for start < end {
		if nums[start] != nums[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func digitsToArray(num int) []int {
	if num == 0 {
		return []int{0}
	}

	if num < 0 {
		num = -num
	}

	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}
	return digits
}
