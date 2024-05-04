package Array

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for _, tt := range tests {
		t.Run("Generate Pascal Triangle", func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res
}
