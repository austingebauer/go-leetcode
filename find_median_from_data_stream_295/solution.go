package find_median_from_data_stream_295

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	v := *h
	p := v[len(v)-1]
	*h = v[0 : len(v)-1]
	return p
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
	v := *h
	p := v[len(v)-1]
	*h = v[0 : len(v)-1]
	return p
}

type MedianFinder struct {
	// keep the smaller half of numbers in a max heap for O(1) access to middle
	smallerHalf *MaxHeap

	// keep the larger half of numbers in a min heap for O(1) access to middle
	largerHalf *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	return MedianFinder{
		largerHalf:  minHeap,
		smallerHalf: maxHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// push num onto the max heap for the smaller half of numbers
	heap.Push(this.smallerHalf, num)

	// to maintain balance between the two heaps, offer
	// largest from smaller half to the larger half
	heap.Push(this.largerHalf, heap.Pop(this.smallerHalf).(int))

	// if the larger half becomes larger in length that the smaller half,
	// we need to offer the lowest in the largest half to the smaller half
	if this.largerHalf.Len() > this.smallerHalf.Len() {
		heap.Push(this.smallerHalf, heap.Pop(this.largerHalf).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smallerHalf.Len() == this.largerHalf.Len() {
		// there is an even number of elements in the MedianFinder, so the median is the
		// average of largest of the smaller half and the smallest of the larger half
		return (float64((*this.smallerHalf)[0]) + float64((*this.largerHalf)[0])) / 2
	} else {
		// there is an odd number of elements in the MedianFinder,
		// so the median is the largest of the smaller half
		return float64((*this.smallerHalf)[0])
	}
}
