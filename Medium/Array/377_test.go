package Array

import (
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "Simple Case", nums: []int{1, 2, 3}, target: 4, want: 7},            // Combinations: {1,1,1,1}, {1,1,2}, {2,2}, {1,3}, {3,1}, {2,1,1}, {1,2,1}
		{name: "Empty Nums", nums: []int{}, target: 0, want: 1},                    // One way to make 0 is to use no numbers
		{name: "No Valid Combos", nums: []int{2, 5}, target: 3, want: 0},           // No way to sum to 3 using 2 and 5
		{name: "Single Element Multiple Uses", nums: []int{3}, target: 9, want: 1}, // Only one way using {3, 3, 3}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum4(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("%s failed: combinationSum4(%v, %d) got %d, want %d", tt.name, tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func combinationSum4(nums []int, target int) int {
	return 0
}
