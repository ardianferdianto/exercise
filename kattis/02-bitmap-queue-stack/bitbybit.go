package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const registerSize = 32

type Register [registerSize]interface{}

func clear(register *Register, bit int) {
	if bit == -1 {
		return
	}
	register[bit] = false
}

func set(register *Register, bit int) {
	if bit == -1 {
		return
	}
	register[bit] = true
}

func orOp(register *Register, dest, src1, src2 int) {
	if src1 == -1 || src2 == -1 {
		register[dest] = nil
	}
	if register[src1] == nil || register[src2] == nil {
		if register[src1].(bool) == true || register[src2].(bool) == true {
			register[dest] = true
		} else {
			register[dest] = nil
		}
	} else {
		register[dest] = register[src1].(bool) || register[src2].(bool)
	}
}

func andOp(register *Register, dest, src1, src2 int) {
	if src1 == -1 || src2 == -1 {
		register[dest] = nil
	}
	if register[src1] == nil || register[src2] == nil {
		if register[src1].(bool) == false || register[src2].(bool) == false {
			register[dest] = false
		} else {
			register[dest] = nil
		}
	} else {
		register[dest] = register[src1].(bool) && register[src2].(bool)
	}
}

func printRegister(register Register) {
	result := make([]string, registerSize)

	for i := 31; i >= 0; i-- {
		if register[i] == nil {
			result[31-i] = "?"
		} else if value, ok := register[i].(bool); ok && value == true {
			result[31-i] = "1"
		} else {
			result[31-i] = "0"
		}
	}
	fmt.Println(strings.Join(result, ""))
}

func processInstructionSequence(scanner *bufio.Scanner) {
	var n int
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &n)

		if n == 0 {
			break
		}

		var register Register
		var sequence []string
		for i := 0; i < n; i++ {
			scanner.Scan()
			sequence = append(sequence, scanner.Text())
		}

		processInstructions(&register, sequence)
	}
}

func processInstructions(register *Register, instructions []string) {
	print := true
	for _, instruction := range instructions {
		var op string
		var dest, src1, src2 int
		fields := strings.Fields(instruction)

		fmt.Sscan(instruction, &op, &dest)
		switch op {
		case "CLEAR":
			if len(fields) < 2 {
				dest = -1
			}
			clear(register, dest)
		case "SET":
			if len(fields) < 2 {
				dest = -1
			}
			set(register, dest)
		case "OR":
			fmt.Sscan(instruction, &op, &dest, &src2)
			src1 = dest
			if len(fields) < 3 {
				src2 = -1
			}
			orOp(register, dest, src1, src2)
		case "AND":
			fmt.Sscan(instruction, &op, &dest, &src2)
			src1 = dest
			if len(fields) < 3 {
				src2 = -1
			}
			andOp(register, dest, dest, src2)
		default:
			print = false
		}
	}
	if len(instructions) == 0 || print == false {
		return
	}
	printRegister(*register)
}
func parseIndex(s string) uint {
	var i uint
	fmt.Sscanf(s, "%d", &i)
	return i
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	processInstructionSequence(scanner)
}
