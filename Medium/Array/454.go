package Array

import (
	"fmt"
)

func FourSumCountTest() {
	testCases := []struct {
		nums1    []int
		nums2    []int
		nums3    []int
		nums4    []int
		expected int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0, 1}, []int{-1, 0}, []int{0, 1}, []int{1, 0}, 17},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
		{[]int{}, []int{}, []int{}, []int{}, 0},
		{[]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}, 6},
	}

	for _, tc := range testCases {
		result := fourSumCount(tc.nums1, tc.nums2, tc.nums3, tc.nums4)
		if result == tc.expected {
			fmt.Printf("Test Passed for nums1: %v, nums2: %v, nums3: %v, nums4: %v, Expected and got %d\n", tc.nums1, tc.nums2, tc.nums3, tc.nums4, result)
		} else {
			fmt.Printf("Test Failed for nums1: %v, nums2: %v, nums3: %v, nums4: %v, Expected %d, Got %d\n", tc.nums1, tc.nums2, tc.nums3, tc.nums4, tc.expected, result)
		}
	}
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	res := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			res[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			count += res[-(nums3[i] + nums4[j])]
		}
	}

	return count
}
