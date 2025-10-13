package main

import (
	"container/heap"
	"fmt"
)

// Top K Frequent Elements
// Approach: Frequency map + Min-Heap

type Pair struct {
	num  int
	freq int
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, Pair{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Pair).num)
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [2 1] or [1 2]
}
