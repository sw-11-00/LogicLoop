package Array

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Test1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "Test2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "Test3",
			nums: []int{},
			want: [][]int{{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for _, num := range nums {
		for _, r := range res {
			res = append(res, append([]int{num}, r...))
		}
	}
	return res
}
