package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 1},
		{[]int{99, 99}, 1, 99},
	}

	for idx, tc := range testCases {
		result := findKthLargest(tc.nums, tc.k)
		if result == tc.expected {
			fmt.Printf("Test Case %d: Passed - Got %d\n", idx, result)
		} else {
			fmt.Printf("Test Case %d: Failed - Got %d, Expected %d\n", idx, result, tc.expected)
		}
	}
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}
