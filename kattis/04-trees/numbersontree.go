package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	var path string
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %s", &n, &path)
	}
	cur := 1
	for _, c := range path {
		cur *= 2
		if c == 'R' {
			cur++
		}
	}
	result := (1 << (n + 1)) - cur
	fmt.Println(result)
}
