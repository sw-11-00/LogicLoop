package Array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{"empty array", []int{}, 0, nil},
		{"single element", []int{1}, 1, nil},
		{"two elements", []int{1, 2}, 3, []int{1, 2}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{"no solution", []int{1, 2, 3, 4, 5}, 10, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	for start, end := 0, len(numbers)-1; start < end; {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum > target {
			end--
		} else if sum < target {
			start++
		}
	}

	return nil
}
