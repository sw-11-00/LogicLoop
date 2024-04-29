package Array

import (
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			name: "Test2",
			nums: []int{7, 6, 5, 4, 3, 2, 1},
			want: 0,
		},
		{
			name: "Test3",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: 6,
		},
		{
			name: "Test4",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Test5",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.nums); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func maxProfit(prices []int) int {
	max := 0
	slow := 0
	fast := 1
	if len(prices) == 1 {
		return 0
	}

	for fast < len(prices) {
		if prices[slow] > prices[fast] {
			slow = fast
		} else {
			max = Max(prices[fast]-prices[slow], max)
		}
		fast++
	}

	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
