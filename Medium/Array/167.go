package Array

import (
	"fmt"
	"reflect"
)

func TwoSumTest() {
	testCases := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3}, 4, []int{1, 3}},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
		{[]int{5, 25, 75}, 100, []int{2, 3}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.numbers, tc.target)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test Passed for numbers %v and target %d: Expected and got %v\n", tc.numbers, tc.target, result)
		} else {
			fmt.Printf("Test Failed for numbers %v and target %d: Expected %v, Got %v\n", tc.numbers, tc.target, tc.expected, result)
		}
	}
}

func twoSum(numbers []int, target int) []int {
	nums := numbers
	numSet := make(map[int]int, len(nums))
	for i, num := range nums {
		numSet[num] = i
	}

	ans := make([]int, 0)
	if numSet[target-nums[0]] > 0 {
		ans = append(ans, 1)
		ans = append(ans, numSet[target-nums[0]]+1)
		return ans
	}
	for i := 1; i < len(nums); i++ {
		if numSet[target-nums[i]] > 0 && numSet[target-nums[i]] != i {
			ans = append(ans, i+1)
			ans = append(ans, numSet[target-nums[i]]+1)
			return ans
		}
	}

	return nil
}
