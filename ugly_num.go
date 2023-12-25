package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// n := 1
	fmt.Printf("result %+v \n", hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	return bits.OnesCount32(uint32(x ^ y))
}
func isUgly(n int) bool {
	u := []int{2, 3, 5}
	if n == 1 {
		return true
	}
	if n < 0 {
		return false
	}
	for _, c := range u {
		for n%c == 0 {
			n /= c
			if n == 1 {
				return true
			}
		}
	}
	return false
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
