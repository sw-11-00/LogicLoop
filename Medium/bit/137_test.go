package bit

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Test1", []int{2, 2, 3, 2}, 3},
		{"Test2", []int{0, 1, 0, 1, 0, 1, 99}, 99},
		{"Test3", []int{-2, -2, 1, -2}, 1},
		{"Test4", []int{1, 1, 1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber3(tt.nums)
			if got != tt.want {
				t.Errorf("%s failed: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func singleNumber(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	for num, count := range freq {
		if count == 1 {
			return num
		}
	}
	return 0
}

func singleNumber3(nums []int) int {
	var result int
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += (num >> i) & 1
		}
		if sum%3 != 0 {
			result |= 1 << i
		}
	}
	if result >= (1 << 31) {
		result -= 1 << 32
	}
	return result
}
