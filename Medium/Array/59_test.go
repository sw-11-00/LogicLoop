package Array

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "test1",
			n:    3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "test2",
			n:    1,
			want: [][]int{
				{1},
			},
		},
		{
			name: "test3",
			n:    4,
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix(%d) got %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func generateMatrix(n int) [][]int {
	left, right, low, high := 0, n-1, 0, n-1
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1

	for num <= n*n {
		for i := left; i <= right; i++ {
			matrix[low][i] = num
			num++
		}
		low++
		for i := low; i <= high; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[high][i] = num
			num++
		}
		high--
		for i := high; i >= low; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
