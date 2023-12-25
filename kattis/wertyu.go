package main

import (
	"bufio"
	"fmt"
	"os"
)

// The main function reads input from a file, prints each line, and then scans and prints a specific
// value from the input.
// PROBLEM : https://open.kattis.com/contests/xoh92u/problems/wertyu
func main() {
	// file, err := os.Open("wertyu/wertyu-1.in")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(convert(scanner.Text()))
	// }
	scanner := bufio.NewScanner(os.Stdin)
	// var input string
	for scanner.Scan() {
		fmt.Println(convert(scanner.Text()))
	}

}

func convert(s string) string {
	qwertyMap := map[string]string{
		"W": "Q", ",": "M",
		"E": "W", ".": ",",
		"R": "E", "/": ".",
		"T": "R", ";": "L",
		"Y": "T", "'": ";",
		"U": "Y", "[": "P",
		"I": "U", "]": "[",
		"O": "I", "\\": "]",
		"P": "O", "=": "-",
		"S": "A", "-": "0",
		"D": "S", "0": "9",
		"F": "D", "9": "8",
		"G": "F", "8": "7",
		"H": "G", "7": "6",
		"J": "H", "6": "5",
		"K": "J", "5": "4",
		"L": "K", "4": "3",
		"X": "Z", "3": "2",
		"C": "X", "2": "1",
		"V": "C", "1": "`",
		"B": "V",
		"N": "B",
		"M": "N",
	}
	res := ""
	for _, char := range s {
		if char == ' ' {
			res += string(char)
			continue
		}
		res += qwertyMap[string(char)]
	}
	return res
}
