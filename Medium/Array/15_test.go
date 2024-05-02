package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Test1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Test2",
			nums: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Test3",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "Test4",
			nums: []int{0},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum1(tt.nums)
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
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

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			for j, k := i+1, len(nums)-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					for j < k && nums[j] == nums[j+1] {
						j++
						continue
					}
					for j < k && nums[k] == nums[k-1] {
						k--
						continue
					}
					j++
					k--
				} else if sum > 0 {
					k--
				} else {
					j++
				}
			}
		}
	}

	return res
}
