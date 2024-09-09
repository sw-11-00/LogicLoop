package Array

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{"Test 1", 3, 7, [][]int{{1, 2, 4}}},
		{"Test 2", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"Test 4", 3, 15, [][]int{{1, 5, 9}, {1, 6, 8}, {2, 4, 9}, {2, 5, 8}, {2, 6, 7}, {3, 4, 8}, {3, 5, 7}, {4, 5, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.k, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3(%d, %d) = %v, want %v", tt.k, tt.n, got, tt.want)
			}
		})
	}
}

func combinationSum3(k int, n int) [][]int {
	var results [][]int
	var backpack func(comb []int, start, target int)
	backpack = func(comb []int, start, target int) {
		if len(comb) == k && target == 0 {
			temp := make([]int, k)
			copy(temp, comb)
			results = append(results, temp)
			return
		}

		for i := start; i <= 9; i++ {
			if target-i >= 0 {
				backpack(append(comb, i), i+1, target-i)
			}
		}
	}

	backpack([]int{}, 1, n)

	return results
}

func combinationSum31(k int, n int) [][]int {
	var res [][]int
	var backtrack func(start int, n int, path []int)

	backtrack = func(start int, n int, path []int) {
		if n == 0 && len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := start; i <= 9; i++ {
			if n-i >= 0 {
				backtrack(i+1, n-i, append(path, i))
			}
		}
	}

	backtrack(1, n, []int{})

	return res
}
