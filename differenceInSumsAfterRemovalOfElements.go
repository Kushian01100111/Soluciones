package main

import (
	"container/heap"
	"math"
)

// MinHeap para suffix
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MaxHeap para prefix (guardamos los negativos para simular un max-heap)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Mayor primero
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumDifference(nums []int) int64 {
	size := len(nums)
	n := size / 3
	prefixSum := make([]int64, size)
	suffixSum := make([]int64, size)

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	var prefix int64

	for i := 0; i < size; i++ {
		heap.Push(maxHeap, nums[i])
		prefix += int64(nums[i])
		if maxHeap.Len() > n {
			removed := heap.Pop(maxHeap).(int)
			prefix -= int64(removed)
		}
		if i >= n-1 {
			prefixSum[i] = prefix
		}
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	var suffix int64

	for j := size - 1; j > 0; j-- {
		heap.Push(minHeap, nums[j])
		suffix += int64(nums[j])
		if minHeap.Len() > n {
			removed := heap.Pop(minHeap).(int)
			suffix -= int64(removed)
		}
		if j <= size-n {
			suffixSum[j] = suffix
		}
	}

	minDiff := int64(math.MaxInt64)
	for i := n - 1; i < size-n; i++ {
		diff := prefixSum[i] - suffixSum[i+1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
