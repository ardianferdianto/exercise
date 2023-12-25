package main

import (
	"fmt"
)

func main() {
	ranks := []int{13, 2, 3, 1, 9}
	suits := []byte("aaaaa")
	ranks = []int{4, 4, 2, 4, 4}
	suits = []byte("daabc")
	ranks = []int{2, 1, 2, 1, 1}
	suits = []byte("dbaad")
	fmt.Println("Results:", bestHand(ranks, suits))

}

func bestHand(ranks []int, suits []byte) string {
	suitsMap := make(map[byte][]int)
	ranksMap := make(map[int][]byte)
	for i, val := range ranks {
		suitsMap[suits[i]] = append(suitsMap[suits[i]], val)
		ranksMap[val] = append(ranksMap[val], suits[i])
	}
	for _, value := range suitsMap {
		count := len(value)
		if count >= 5 {
			return "Flush"
		}
	}
	result := "High Card"
	for _, value := range ranksMap {
		count := len(value)
		if count >= 3 {
			return "Three of a Kind"
		}
		if count >= 2 {
			result = "Pair"
		}
	}

	return result
}
