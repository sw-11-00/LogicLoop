package dp

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "Test2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "Test3",
			nums: []int{-2, 3, -4},
			want: 24,
		},
		{
			name: "Test4",
			nums: []int{-2},
			want: -2,
		},
		{
			name: "Test5",
			nums: []int{0, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	globalMax := nums[0]
	localMax := nums[0]
	localMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			localMax, localMin = localMin, localMax
		}

		localMax = max(localMax*nums[i], nums[i])
		localMin = min(localMin*nums[i], nums[i])
		globalMax = max(globalMax, localMax)
	}

	return globalMax
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
