package main

import (
	"container/heap"
	"log"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Pop() interface{} {
	old := *h
	length := len(old)
	x := old[length-1:]
	*h = old[:length-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func findKthLargest(nums []int, k int) int {
	intHeap := &IntHeap{}

	for idx := range nums {
		*intHeap = append(*intHeap , nums[idx])
	}
	heap.Init(intHeap)

	heap.Push(intHeap , k)
	log.Println(intHeap)
	return intHeap.Pop().(int)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 3
	log.Println(findKthLargest(nums, k))
}
