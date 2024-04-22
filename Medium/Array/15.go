package Array

import (
	"fmt"
	"reflect"
	"sort"
)

func ThreeSumTest() {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{3, -2, 1, 0}, nil},
		{[]int{}, nil},
		{[]int{1, 2, -2, -1}, nil},
	}

	for _, tc := range testCases {
		result := threeSum(tc.nums)
		if reflect.DeepEqual(normalize(result), normalize(tc.expected)) {
			fmt.Printf("Test Passed for %v: Expected and got %v\n", tc.nums, result)
		} else {
			fmt.Printf("Test Failed for %v: Expected %v, Got %v\n", tc.nums, tc.expected, result)
		}
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			s := nums[i] + nums[j] + nums[k]
			if s < 0 {
				j++
			} else if s > 0 {
				k--
			} else {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return results
}

func normalize(triplets [][]int) [][]int {
	for _, t := range triplets {
		sort.Ints(t)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for x := range triplets[i] {
			if triplets[i][x] != triplets[j][x] {
				return triplets[i][x] < triplets[j][x]
			}
		}
		return false
	})
	return triplets
}
