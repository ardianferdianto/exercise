package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	var result int
	if scanner.Scan() {
		fmt.Sscan(scanner.Text(), &n)
	}
	result = n
	hail(n, &result)
	fmt.Println(result)
}

func hail(n int, result *int) int {
	if n == 1 {
		return n
	}
	if n%2 == 0 {
		*result += (n / 2)
		return hail(n/2, result)
	}
	*result += ((3 * n) + 1)
	return hail(((3 * n) + 1), result)
}
