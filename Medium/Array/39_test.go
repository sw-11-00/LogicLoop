package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "test1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "test2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "test3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{}, // No combination
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.candidates, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum(%v, %d) got %v, want %v", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int

	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
			backtrack(i, target-candidates[i], append(path, candidates[i]))
		}
	}

	backtrack(0, target, []int{})

	return res
}
