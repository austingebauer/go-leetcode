package top_k_frequent_elements_347

import "container/heap"

type freqHeapVal struct {
	key int // the identifier
	val int // number of times key occurred
}

type freqHeap []freqHeapVal

func (h freqHeap) Len() int {
	return len(h)
}
func (h freqHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h freqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *freqHeap) Push(f interface{}) {
	*h = append(*h, f.(freqHeapVal))
}
func (h *freqHeap) Pop() interface{} {
	heapV := *h
	p := heapV[len(heapV)-1]
	*h = heapV[0 : len(heapV)-1]
	return p
}

// Time:  O(n)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	// build map of frequencies
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n] += 1
	}

	// push frequencies into freqHeap
	h := &freqHeap{}
	heap.Init(h)
	for k, v := range freq {
		heap.Push(h, freqHeapVal{
			key: k,
			val: v,
		})
	}

	// pop k largest frequency items from the freqHeap
	top := make([]int, k)
	for i := 0; i < k; i++ {
		top[i] = heap.Pop(h).(freqHeapVal).key
	}

	return top
}
