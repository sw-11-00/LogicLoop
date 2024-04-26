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
		{2, []int{1, 1, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{5, []int{1, 5, 10, 10, 5, 1}},
	}

	for _, tt := range tests {
		t.Run("GetRow Pascal Triangle", func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}

func getRow(rowIndex int) []int {
	ans := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans[rowIndex]
}
