package Array

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "Test1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "Test2",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "Test3",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge1(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums0 := make([]int, len(nums1))
	i := 0
	j := 0
	k := 0
	for ; k < len(nums1) && i < m && j < n; k++ {
		if nums1[i] < nums2[j] {
			nums0[k] = nums1[i]
			i++
		} else {
			nums0[k] = nums2[j]
			j++
		}
	}

	if i < m {
		nums0 = append(nums0[:k], nums1[i:]...)
	}
	if j < n {
		nums0 = append(nums0[:k], nums2[j:]...)
	}

	copy(nums1, nums0)
}
