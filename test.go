package main

import (
	"fmt"
)

func main() {
	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c = []int{1, 3, 7, 11, 12, 14, 18}

	c = []int{1, 2}
	s := "ba"
	// c = []int{12, 23, 36, 46, 62}
	// s = "spuda"
	fmt.Println("result ", string(slowestKey(c, s)))

}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	length := len(releaseTimes)
	b := []byte(keysPressed)
	elapsed := 0
	max := 0
	longestKey := 0
	before := 0
	for i := 0; i < length; i++ {
		if i-1 < 0 {
			before = 0
		} else {
			before = releaseTimes[i-1]
		}
		elapsed = releaseTimes[i] - before
		if elapsed < 0 {
			elapsed = -elapsed
		}
		// fmt.Println(i, max, elapsed)
		if max <= elapsed {
			if max == elapsed && b[longestKey] > b[i] {
				continue
			}
			max = elapsed
			longestKey = i
		}
	}

	return b[longestKey]
}

func lenLongestFibSubseq(arr []int) int {
	// res := []int{}
	// sub := []int{}
	// count := 0
	// max := 0
	memo := make(map[[2]int]int)
	return fibonachi(0, 1, arr, memo)
}

func fibonachi(i int, j int, arr []int, mem map[[2]int]int) int {
	if val, ok := mem[[2]int{i, j}]; ok {
		return val
	}

	size := len(arr)
	if j >= size || i >= size {
		return 0
	}

	nextNum := arr[i] + arr[j]
	count := 0
	nextIndex := -1
	idx1 := 0
	idx2 := 0

	for k := j + 1; k < size; k++ {
		nextNum = arr[k]
		fmt.Printf("%+v\n", nextNum)
		fmt.Println("===")
		// 1, 3, 7, 11, 12, 14, 18
		for p1 := 0; p1 < k-1; p1++ {
			for p2 := p1 + 1; p2 < k; p2++ {
				if p1 == p2 {
					continue
				}
				fmt.Println(p1, p2)
				if arr[p1]+arr[p2] == nextNum {
					nextIndex = k
					idx1 = p1
					idx2 = p2
					fmt.Println("=> ", p1, p2, " = ", nextNum, " nextIndex = ", nextIndex)
					break
				}
			}
		}
	}

	if nextIndex != -1 {
		count = 1 + fibonachi(j, nextIndex, arr, mem)
	}

	mem[[2]int{idx1, idx2}] = count
	// fmt.Println("mem => ", mem[[2]int{idx1, idx2}])

	return count
}
