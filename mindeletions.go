package main

import (
	"fmt"
)

func main() {
	n := "example"
	fmt.Println("s => ", n)
	fmt.Println(minDeletions(n))
}

func minDeletions(s string) int {
	groupedElements := make(map[rune]int)
	freqSet := map[int]bool{}
	delete := 0
	for _, c := range s {
		groupedElements[c]++
	}

	for _, i := range groupedElements {
		if _, ok := freqSet[i]; !ok {
			freqSet[i] = true
		} else {
			j := i
			for ok := freqSet[j]; ok; j-- {
				if j <= 0 {
					break
				}
				if _, ok := freqSet[j]; !ok {
					freqSet[j] = true
					break
				}
				delete++
			}

		}
	}
	return delete
}
