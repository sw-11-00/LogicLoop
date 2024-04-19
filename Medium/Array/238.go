package Array

import (
	"fmt"
	"reflect"
)

func ProductExceptSelfTest() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{[]int{1}, []int{1}}, // Edge case with a single element
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc.nums)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test Passed for %v: Got %v\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %v, Got %v\n", tc.nums, tc.expected, result)
		}
	}
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums)) // 创建一个与 nums 等长的新数组

	// 将新数组的所有元素初始化为1
	for i := range ans {
		ans[i] = 1
	}

	beforeSum := 1
	afterSum := 1
	for i, j := 0, len(nums)-1; i < len(nums); {
		ans[i] *= beforeSum
		ans[j] *= afterSum
		beforeSum *= nums[i]
		afterSum *= nums[j]
		i++
		j--
	}

	return ans
}
