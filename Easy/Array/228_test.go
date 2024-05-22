package Array

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{nums: []int{0, 1, 2, 4, 5, 7}, want: []string{"0->2", "4->5", "7"}},
		{nums: []int{0, 2, 3, 4, 6, 8, 9}, want: []string{"0", "2->4", "6", "8->9"}},
		{nums: []int{}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run("TestSummaryRanges", func(t *testing.T) {
			got := summaryRanges(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("summaryRanges(%v) = %v, want %v", tt.nums, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("summaryRanges(%v) = %v, want %v", tt.nums, got, tt.want)
					return
				}
			}
		})
	}
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var res []string
	for i := 0; i < len(nums); i++ {
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}

		if j == i {
			res = append(res, fmt.Sprint(nums[i]))
		} else {
			res = append(res, fmt.Sprint(nums[i])+"->"+fmt.Sprint(nums[j]))
		}
		i = j
	}
	return res
}
