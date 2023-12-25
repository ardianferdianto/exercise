package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("addingwords/sample.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// tampungan
	variables := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "clear" {
			// clear
			variables = make(map[string]int)
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 3 && parts[0] == "def" {
			// mapping variable
			variableName := parts[1]
			variableValue, _ := strconv.Atoi(parts[2])

			if val, ok := variables[variableName]; ok {
				if val == variableValue {
					delete(variables, variableName)
				}
			} else {
				variables[variableName] = variableValue
			}
		} else if len(parts) >= 4 && parts[0] == "calc" {
			// calc foo + bar =
			var result int
			invalid := false

			for i := 1; i < len(parts); i += 2 {
				operation := parts[i-1]
				operand := parts[i]
				var operandValue int

				if _, ok := variables[operand]; !ok {
					invalid = true
					break
				} else {
					operandValue = variables[operand]
				}
				//fmt.Println(operandValue, operation, variables)

				switch operation {
				case "+":
					result += operandValue
				case "-":
					result -= operandValue
				default:
					result = operandValue
				}
			}
			re := checkForValue(result, variables)
			if invalid {
				re = "unknown"
			}
			fmt.Printf("%s %s\n", strings.Join(parts[1:], " "), re)
		}
	}

}

func checkForValue(valueExpected int, variables map[string]int) string {
	for key, value := range variables {
		if value == valueExpected {
			return key
		}
	}
	return "unknown"
}
