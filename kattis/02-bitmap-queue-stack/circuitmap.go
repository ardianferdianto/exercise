package main

import (
	"bufio"
	"os"
	"strconv"
)

// The main function reads input from a file, prints each line, and then scans and prints a specific
// value from the input.
// PROBLEM : https://open.kattis.com/contests/xoh92u/problems/wertyu
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 0
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		// line := scanner.Text()
		// parts := strings.Fields(line)
		// fmt.Println(n)
	}

	// var input string
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	parts := strings.Fields(line)
	// 	fmt.Println(n)
	// }

}
