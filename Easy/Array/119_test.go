package Array

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		want     []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
	}

	for _, tt := range tests {
		t.Run("Test GetRow", func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res[rowIndex]
}
