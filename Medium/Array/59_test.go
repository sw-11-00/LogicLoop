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
	l, r, t, b := 0, n-1, 0, n-1
	num := 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for num <= n*n {
		for i := l; i <= r; i++ {
			res[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			res[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			res[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			res[i][l] = num
			num++
		}
		l++
	}

	return res
}
