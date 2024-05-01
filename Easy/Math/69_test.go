package Math

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "Test1",
			x:    4,
			want: 2,
		},
		{
			name: "Test2",
			x:    8,
			want: 2,
		},
		{
			name: "Test3",
			x:    0,
			want: 0,
		},
		{
			name: "Test4",
			x:    1,
			want: 1,
		},
		{
			name: "Test5",
			x:    16,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt1(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left - 1
}

func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		} else {
			return mid
		}
	}

	// 这里为什么是 right
	return right
}
