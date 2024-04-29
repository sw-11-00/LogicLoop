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
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	nums0 := make([]int, 0, len(nums1))
	i := 0
	j := 0
	for k := 0; k < len(nums1) && i < m && j < n; k++ {
		if nums1[i] <= nums2[j] {
			nums0 = append(nums0, nums1[i])
			i++
		} else {
			nums0 = append(nums0, nums2[j])
			j++
		}
	}

	if i == m {
		nums0 = append(nums0, nums2[j:]...)
	}
	if j == n {
		nums0 = append(nums0, nums1[i:]...)
	}

	copy(nums1, nums0)
}
