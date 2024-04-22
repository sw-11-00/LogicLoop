package Array

import (
	"fmt"
	"reflect"
	"sort"
)

func FourSumTest() {
	testCases := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{-3, -1, 0, 2, 4, 5}, 0, [][]int{{-3, -1, 0, 4}}},
		{[]int{0, 0, 0, 0}, 1, nil},
		{[]int{0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
		{[]int{-5, -4, -3, -2, 1, 3, 3, 5}, 0, [][]int{{-5, -4, 3, 6}, {-5, -3, 3, 5}}},
	}

	for _, tc := range testCases {
		result := fourSum(tc.nums, tc.target)
		if reflect.DeepEqual(normalize(result), normalize(tc.expected)) {
			fmt.Printf("Test Passed for nums %v and target %d: Expected and got %v\n", tc.nums, tc.target, result)
		} else {
			fmt.Printf("Test Failed for nums %v and target %d: Expected %v, Got %v\n", tc.nums, tc.target, tc.expected, result)
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[i+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			for k, m := j+1, n-1; k < m; {
				if nums[i]+nums[j]+nums[k]+nums[m] < target {
					k++
				} else if nums[i]+nums[j]+nums[k]+nums[m] > target {
					m--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
					k++
					m--
					for k < m && nums[k] == nums[k-1] {
						k++
					}
					for k < m && nums[m] == nums[m+1] {
						m--
					}
				}
			}
		}
	}

	return res
}
