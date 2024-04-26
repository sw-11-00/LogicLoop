package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{[]int{2, 3, 6, 7}, 7, [][]int{{7}}},
	}

	for _, tt := range tests {
		t.Run("CombinationSum2", func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			// Since the order of the combinations and the order of numbers within each combination may vary, we need to sort them before comparing
			sortSlices(got)
			sortSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2(%v, %d) = %v, want %v", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}

func sortSlices(s [][]int) {
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

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var backpack func(comb []int, start, target int)
	backpack = func(comb []int, start, target int) {
		if target == 0 {
			temp := make([]int, len(comb))
			copy(temp, comb)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			backpack(append(comb, candidates[i]), i+1, target-candidates[i])
		}
	}

	backpack([]int{}, 0, target)

	return result
}
