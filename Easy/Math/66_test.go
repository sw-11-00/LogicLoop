package Math

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "Test1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "Test2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "Test3",
			digits: []int{0},
			want:   []int{1},
		},
		{
			name:   "Test4",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "Test5",
			digits: []int{9, 9, 9, 9},
			want:   []int{1, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}

	return append([]int{1}, digits...)
}
