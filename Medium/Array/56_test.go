package Array

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "Test1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Test2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "Test3",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "Test4",
			intervals: [][]int{{1, 4}},
			want:      [][]int{{1, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge1(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
