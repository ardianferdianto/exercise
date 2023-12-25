package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func evaluateCircuit(inputs map[string]bool, circuitDescription string) bool {
	stack := []bool{}

	for _, char := range circuitDescription {
		if char >= 'A' && char <= 'Z' {
			// If the character is an uppercase letter, it's an input variable.
			stack = append(stack, inputs[string(char)])
		} else if char == '*' {
			// AND operation
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1 && op2)
		} else if char == '+' {
			// OR operation
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1 || op2)
		} else if char == '-' {
			// NOT operation
			op := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, !op)
		}
		fmt.Println(stack)
	}

	return stack[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the number of input variables
	scanner.Scan()

	// Read the input truth values
	scanner.Scan()
	inputValues := strings.Fields(scanner.Text())
	inputs := make(map[string]bool)
	for i, value := range inputValues {
		inputs[string('A'+i)] = value == "T"
	}

	// Read the circuit description
	scanner.Scan()
	circuitDescription := scanner.Text()

	// Evaluate the circuit and print the result
	result := evaluateCircuit(inputs, circuitDescription)
	if result {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	n := 0

// 	stack := Stack{}
// 	if scanner.Scan() {
// 		//size, _ := scanner.Text()
// 	}
// 	if scanner.Scan() {
// 		line := scanner.Text()
// 		parts := strings.Fields(line)
// 		for idx, p := range parts {
// 			if p == "T" {
// 				n |= 1 << (idx)
// 			}
// 		}
// 	}

// 	if scanner.Scan() {
// 		line := scanner.Text()
// 		parts := strings.Fields(line)

// 		for _, p := range parts {
// 			if isAlphabet(p[0]) == true {
// 				alphaIndex, _ := getAlphabetOrder(p)
// 				// bit start from 0
// 				if (n & (1 << (alphaIndex - 1))) != 0 {
// 					stack.Push(1)
// 				} else {
// 					stack.Push(0)
// 				}
// 			}

// 			if p == "*" {
// 				stack.Push(stack.Pop() * stack.Pop())
// 			}
// 			if p == "+" {
// 				stack.Push(stack.Pop() + stack.Pop())
// 			}
// 			if p == "-" {
// 				stack.Push(stack.Pop() ^ 1)
// 			}
// 		}

// 		if stack.Pop() <= 0 {
// 			fmt.Print("F")
// 		} else {
// 			fmt.Print("T")
// 		}
// 	}
// }

// func getAlphabetOrder(alphabet string) (int, error) {
// 	// lower case terus pake ascii
// 	alphabet = strings.ToLower(alphabet)
// 	if len(alphabet) != 1 || !('a' <= alphabet[0] && alphabet[0] <= 'z') {
// 		return 0, fmt.Errorf("salah lur", alphabet)
// 	}
// 	// itung pake ascii + 1 => karena ascii start from 0
// 	order := int(alphabet[0] - 'a' + 1)
// 	return order, nil
// }

// func isNotSymbol(s string) bool {
// 	return s != "*" && s != "+" && s != "-"
// }

// func isAlphabet(s byte) bool {
// 	return s >= 'A' && s <= 'Z'
// }

// // Stack represents a stack data structure.
// type Stack struct {
// 	items []int
// }

// func (s *Stack) Push(item int) {
// 	s.items = append(s.items, item)
// }

// func (s *Stack) Pop() int {
// 	if len(s.items) == 0 {
// 		return 0
// 	}

// 	top := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return top
// }

// func (s *Stack) Peek() int {
// 	if len(s.items) == 0 {
// 		return 0
// 	}
// 	return s.items[len(s.items)-1]
// }
