package heap

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{"empty array", []int{}, 0, nil},
		{"single element", []int{1}, 1, []int{1}},
		{"multiple elements", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

type Element struct {
	num       int
	frequency int
}

type MinHeap1 []Element

func (h MinHeap1) Len() int           { return len(h) }
func (h MinHeap1) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap1) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	h := &MinHeap1{}
	heap.Init(h)
	for num, frequency := range frequencyMap {
		if h.Len() < k {
			heap.Push(h, Element{num, frequency})
		} else if frequency > (*h)[0].frequency {
			heap.Pop(h)
			heap.Push(h, Element{num, frequency})
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).(Element).num
	}

	return result
}
