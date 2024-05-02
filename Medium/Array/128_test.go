package Array

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "Test2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "Test3",
			nums: []int{1, 2, 0, 1},
			want: 3,
		},
		{
			name: "Test4",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Test5",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	count := 0
	for num := range numSet {
		if !numSet[num-1] {
			localCount := 1
			j := num
			for numSet[j+1] {
				j++
				localCount++
			}
			count = max(localCount, count)
		}
	}

	return count
}
