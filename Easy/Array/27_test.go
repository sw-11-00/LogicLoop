package Array

import (
	"testing"
)

func TestRemoveElementTest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{
			name: "Test1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			name: "Test2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
		{
			name: "Test3",
			nums: []int{1},
			val:  1,
			want: 0,
		},
		{
			name: "Test4",
			nums: []int{3, 3},
			val:  3,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement3(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func removeElement3(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}

	return i
}
