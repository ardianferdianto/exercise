package main

import (
	"container/heap"
	"fmt"
)

type Word struct {
	word  string
	count int
}

type MHeap []Word

func (h MHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].word < h[j].word
	}

	return h[i].count > h[j].count
}

// len
func (h MHeap) Len() int { return len(h) }

// swap
func (h MHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// push
func (h *MHeap) Push(x interface{}) { *h = append(*h, x.(Word)) }

// pop
func (h *MHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Problems: https://leetcode.com/problems/top-k-frequent-words/solutions/2464732/go-solution-using-hashmap-priorityqueue/
func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}

func topKFrequent(words []string, k int) []string {
	mapWord := make(map[string]int)
	maxHeap := &MHeap{}

	for _, word := range words {
		mapWord[word]++
	}

	for k, v := range mapWord {
		heap.Push(maxHeap, Word{word: k, count: v})
	}

	var result []string
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(maxHeap).(Word).word)
	}

	return result
}
