package Array

import (
	"fmt"
	"reflect"
)

func TwoSumTest() {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		//{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		//{[]int{3, 2, 4}, 6, []int{1, 2}},
		//{[]int{3, 3}, 6, []int{0, 1}},
		//{[]int{1, 2, 3}, 7, nil},
		//{[]int{4, 2, 5, 1}, 8, nil},
		{[]int{1, 3, 4, 2}, 6, []int{2, 3}},
		//{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test Passed for nums %v and target %d: Expected and got %v\n", tc.nums, tc.target, result)
		} else {
			fmt.Printf("Test Failed for nums %v and target %d: Expected %v, Got %v\n", tc.nums, tc.target, tc.expected, result)
		}
	}
}

func twoSum(nums []int, target int) []int {
	numSet := make(map[int]int, len(nums))
	for i, num := range nums {
		numSet[num] = i
	}

	ans := make([]int, 0)
	if numSet[target-nums[0]] > 0 {
		ans = append(ans, 0)
		ans = append(ans, numSet[target-nums[0]])
		return ans
	}
	for i := 1; i < len(nums); i++ {
		if numSet[target-nums[i]] > 0 && numSet[target-nums[i]] != i {
			ans = append(ans, i)
			ans = append(ans, numSet[target-nums[i]])
			return ans
		}
	}

	return nil
}
