package Array

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "Test1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "Test2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "Test3",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			want:        [][]int{{5, 7}},
		},
		{
			name:        "Test4",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
		{
			name:        "Test5",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			want:        [][]int{{1, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.intervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
