package main

import (
	"bufio"
	"fmt"
	"os"
)

// lemot fak
func bellTollCount(creatures []rune) int {
	bellTolls := 0

	// keep running as long as contain ocelot
	for containsOcelot(creatures) {
		bellTolls++

		// Find ocelot lowest in the pile turns into a zebra
		index := indexOfOcelot(creatures)
		if index < 0 {
			break
		}
		creatures[index] = 'Z'

		// Turn all zebras below the lowest ocelot into ocelot
		for i := index + 1; i < len(creatures) && creatures[i] == 'Z'; i++ {
			creatures[i] = 'O'
		}

	}

	return bellTolls
}

func bellTollCountV2(creatures []rune) int {
	bellTolls := 0
	lastOcelotIndex := -1

	for i := len(creatures) - 1; i >= 0; i-- {
		if creatures[i] == 'O' {
			lastOcelotIndex = i
			break
		}
	}

	for lastOcelotIndex >= 0 {
		bellTolls++
		creatures[lastOcelotIndex] = 'Z'

		for i := lastOcelotIndex + 1; i < len(creatures) && creatures[i] == 'Z'; i++ {
			creatures[i] = 'O'
		}

		lastOcelotIndex = -1
		for i := len(creatures) - 1; i >= 0; i-- {
			if creatures[i] == 'O' {
				lastOcelotIndex = i
				break
			}
		}
	}

	return bellTolls
}

func containsOcelot(creatures []rune) bool {
	for _, c := range creatures {
		if c == 'O' {
			return true
		}
	}
	return false
}

func indexOfOcelot(creatures []rune) int {
	for i := len(creatures) - 1; i >= 0; i-- {
		if creatures[i] == 'O' {
			return i
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	if scanner.Scan() {
		fmt.Sscan(scanner.Text(), &n)
	}

	var creatures []rune
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) > 0 {
			creatures = append(creatures, []rune(line)[0])
		}
	}

	// Edge case: No creatures or all zebras
	if n == 0 || !containsOcelot(creatures) {
		fmt.Println(0)
		return
	}

	output := bellTollCount(creatures)
	fmt.Println(output)
}
