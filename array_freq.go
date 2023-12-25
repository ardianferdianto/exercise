package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := []int{1, 1, 2, 2, 2, 3}
	fmt.Printf("result %+v\n", frequencySort2(mat))
}

func frequencySort2(nums []int) []int {
	groupedElements := make(map[int]int)
	for _, num := range nums {
		groupedElements[num]++
	}
	var pairs [][2]int
	for i, value := range groupedElements {
		pairs = append(pairs, [2]int{i, value})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] > pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	sortedSlice := []int{}
	i := 0
	for len(pairs) > 0 {
		pair := pairs[i]
		pairs = pairs[i+1:]
		for j := 0; j < pair[1]; j++ {
			sortedSlice = append(sortedSlice, pair[0])
		}
	}
	return sortedSlice
}

func frequencySort(nums []int) []int {
	// sort.Ints(nums)
	groupedElements := make(map[int]int)

	for _, num := range nums {
		groupedElements[num]++
	}
	// fmt.Println(groupedElements)
	var pairs []struct {
		Key   int
		Value int
	}
	for key, value := range groupedElements {
		pairs = append(pairs, struct {
			Key   int
			Value int
		}{Key: key, Value: value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Key > pairs[j].Key
		}
		return pairs[i].Value < pairs[j].Value
	})

	sortedSlice := []int{}
	for _, pair := range pairs {
		for j := 0; j < pair.Value; j++ {
			sortedSlice = append(sortedSlice, pair.Key)
		}
	}

	return sortedSlice
}
