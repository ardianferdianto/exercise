package main

import (
	"fmt"
	"strings"
)

func getAlphabetOrder(alphabet string) (int, error) {
	// Convert the alphabet to lowercase for case-insensitivity
	alphabet = strings.ToLower(alphabet)

	// Check if the input is a single alphabet character
	if len(alphabet) != 1 || !('a' <= alphabet[0] && alphabet[0] <= 'z') {
		return 0, fmt.Errorf("invalid input: %s is not a single alphabet character", alphabet)
	}

	// Calculate the order number
	order := int(alphabet[0] - 'a' + 1)
	return order, nil
}

func main() {
	// Get input alphabet from the user
	var input string
	fmt.Print("Enter an alphabet: ")
	fmt.Scan(&input)

	// Get the order number
	order, err := getAlphabetOrder(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Printf("The order number of %s is: %d\n", input, order)
}
