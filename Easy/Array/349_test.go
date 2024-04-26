package Array

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "test1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2},
		},
		{
			name:  "test2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{9, 4},
		},
		{
			name:  "test3",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, len(nums1)+len(nums2))
	numSet := make(map[int]bool, len(nums1)+len(nums2))
	ans := make([]int, 0)
	for _, num := range nums1 {
		numMap[num] += 1
		numSet[num] = true
	}

	for _, num := range nums2 {
		if numMap[num] > 0 && numSet[num] == true {
			ans = append(ans, num)
			numSet[num] = false
		}
	}

	return ans
}
