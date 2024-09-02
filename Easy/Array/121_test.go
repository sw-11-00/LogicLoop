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
	minPrice := prices[0]
	mxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		mxProfit = max(mxProfit, prices[i]-minPrice)
	}

	return mxProfit
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
