package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value  int
	Parent *Node
}

func main() {
	nodes := make([]*Node, 101)
	// read buffer
	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i < 101; i++ {
		nodes[i] = &Node{Value: i}
		println(i, nodes[i].Value)
	}

	var k int
	if scanner.Scan() {
		fmt.Sscan(scanner.Text(), &k)
	}

	current := nodes[k]

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		parentValue := toInt(fields[0])

		if parentValue == -1 {
			break
		}

		for i := 1; i < len(fields); i++ {
			childValue := toInt(fields[i])
			nodes[childValue].Parent = nodes[parentValue]
		}
	}

	for current != nil {
		fmt.Print(current.Value)
		fmt.Print(" ")
		current = current.Parent
	}
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// func mainold() {
// 	// prepare
// 	nodes := make(map[int]*Node)
// 	// read buffer
// 	scanner := bufio.NewScanner(os.Stdin)

// 	var k int
// 	if scanner.Scan() {
// 		fmt.Sscan(scanner.Text(), &k)
// 	}
// 	scanner.Scan()
// 	rootValue, _ := strconv.Atoi(strings.Fields(scanner.Text())[1])
// 	root := &Node{Value: rootValue}
// 	nodes[rootValue] = root

// 	// assign to tree
// 	for scanner.Scan() {
// 		fields := strings.Fields(scanner.Text())
// 		parentValue, _ := strconv.Atoi(fields[0])

// 		if parentValue == -1 {
// 			break
// 		}

// 		parentNode, ok := nodes[parentValue]
// 		if !ok {
// 			parentNode = &Node{Value: parentValue}
// 			nodes[parentValue] = parentNode
// 		}

// 		for i := 1; i < len(fields); i++ {
// 			childValue, _ := strconv.Atoi(fields[i])
// 			childNode, ok := nodes[childValue]
// 			if !ok {
// 				childNode = &Node{Value: childValue}
// 				nodes[childValue] = childNode
// 			}
// 			parentNode.Children = append(parentNode.Children, childNode)
// 		}
// 	}
// 	// traverse tree
// 	printTree(root, 0)
// }

// func printTree(node *Node, level int) {
// 	fmt.Printf("%s%d\n", strings.Repeat("  ", level), node.Value)
// 	for _, child := range node.Children {
// 		// fmt.Println("Child found")
// 		printTree(child, level+1)
// 	}
// }
