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
		{"Test 3", 2, 18, [][]int{}}, // No possible combinations
		{"Test 4", 3, 15, [][]int{{1, 5, 9}, {2, 4, 9}, {2, 5, 8}, {3, 4, 8}, {3, 5, 7}, {4, 5, 6}}},
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
			if target-i >= 0 && len(comb) < k {
				backpack(append(comb, i), i+1, target-i)
			}
		}
	}

	backpack([]int{}, 1, n)

	return results
}
