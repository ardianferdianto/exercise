package main

import (
	"container/heap"
	"fmt"
)

// 215. Kth Largest Element in an Array
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	fmt.Println(findKthLargest(nums, k))
}

// no sorting
func findKthLargest(nums []int, k int) int {
	maxHeap := &MaxHeap{}
	*maxHeap = nums
	heap.Init(maxHeap)
	res := 0

	for i := 0; i < k; i++ {
		res = heap.Pop(maxHeap).(int)
	}
	return res
}

type MaxHeap []int

func (maxHeap MaxHeap) Len() int            { return len(maxHeap) }
func (maxHeap MaxHeap) Less(i, j int) bool  { return maxHeap[i] > maxHeap[j] }
func (maxHeap MaxHeap) Swap(i, j int)       { maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i] }
func (maxHeap *MaxHeap) Push(x interface{}) { *maxHeap = append(*maxHeap, x.(int)) }
func (maxHeap *MaxHeap) Pop() interface{} {
	old := *maxHeap
	n := len(*maxHeap)
	top := old[n-1]
	*maxHeap = old[:n-1]
	return top
}

func findKthLargestArrayHeap(nums []int, k int) int {
	for i := len(nums)/2 - 1; i > -1; i-- {
		heapify(nums, 0, len(nums)-1, i)
	}

	var res int
	lastIdx := len(nums) - 1

	for i := 0; i < k; i++ {
		res = nums[0]
		nums[0], nums[lastIdx] = nums[lastIdx], nums[0]
		heapify(nums, 0, lastIdx-1, 0)
		lastIdx--
	}

	return res
}

func heapify(nums []int, low int, high int, parent int) {
	left := 2*parent + 1
	right := 2*parent + 2

	larger := parent

	if low <= left && left <= high {
		if nums[larger] < nums[left] {
			larger = left
		}
	}

	if low <= right && right <= high {
		if nums[larger] < nums[right] {
			larger = right
		}
	}

	if larger != parent {
		nums[larger], nums[parent] = nums[parent], nums[larger]
		heapify(nums, low, high, larger)
	}
}
