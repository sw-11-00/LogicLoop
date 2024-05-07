package bit

import (
	"reflect"
	"testing"
)

func TestSingleNumber1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"Test1", []int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{"Test2", []int{2, 2, 1, 1, 3, 5}, []int{3, 5}},
		{"Test3", []int{1, 2, 1, 2, 3, 5}, []int{3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber2(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func singleNumber1(nums []int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	var res []int
	for num, count := range freq {
		if count == 1 {
			res = append(res, num)
		}
	}
	return res
}

func singleNumber2(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	mask := 1
	for (mask & xor) == 0 {
		mask <<= 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if (mask & num) == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
