package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{"less than four elements", []int{1, 2, 3}, 6, [][]int{}},
		{"four elements", []int{1, 2, 3, 4}, 10, [][]int{{1, 2, 3, 4}}},
		{"multiple elements", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"no solution", []int{1, 2, 3, 4, 5}, 100, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fourSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
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
