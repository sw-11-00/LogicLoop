package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{n: 4, k: 2, want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{n: 1, k: 1, want: [][]int{{1}}},
		{n: 5, k: 3, want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}},
	}

	for _, tt := range tests {
		t.Run("Combine", func(t *testing.T) {
			got := combine(tt.n, tt.k)
			sortSliceOfSlices(got)
			sortSliceOfSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}

func sortSliceOfSlices(s [][]int) {
	sort.Slice(s, func(i, j int) bool {
		for x := range s[i] {
			if x >= len(s[j]) {
				return false
			}
			if s[i][x] != s[j][x] {
				return s[i][x] < s[j][x]
			}
		}
		return len(s[i]) < len(s[j])
	})
	for _, arr := range s {
		sort.Ints(arr)
	}
}

func combine(n int, k int) [][]int {
	var results [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			results = append(results, temp)
			return
		}

		for i := start; i <= n; i++ {
			backtrack(i+1, append(path, i))
		}
	}

	backtrack(1, []int{})
	return results
}
