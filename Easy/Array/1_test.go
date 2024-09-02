package Array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Test1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Test2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Test3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Test4",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "Test5",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum1(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func twoSum(nums []int, target int) []int {
	numSet := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numSet[target-num]; ok {
			return []int{i, j}
		}
		numSet[num] = i
	}

	return nil
}

func twoSum1(nums []int, target int) []int {
	twoSumMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := twoSumMap[target-num]; ok {
			return []int{i, j}
		}
		twoSumMap[num] = i
	}

	return nil
}
