package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("conformity/split-00.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	var input string
	scanner.Scan()
	fmt.Sscanln(scanner.Text(), &input)
	fmt.Print(input)

}

// func conformity(s string) string {

// }
