package main

import (
	"fmt"
)

func main() {
	c := [][]int{{1, 0}, {0, 1}}
	// c = [][]int{{1, 0}, {1, 1}}
	// c = [][]int{
	// 	{1, 1, 0, 0},
	// 	{0, 0, 1, 0},
	// 	{0, 0, 1, 0},
	// 	{0, 0, 0, 1}}
	c = [][]int{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}}
	// c = [][]int{
	// 	{1, 0, 0, 1, 0},
	// 	{0, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 0}}
	// c = [][]int{{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 1}, {0, 0, 1, 1}, {0, 0, 0, 1}}
	// c = [][]int{
	// 	{0, 0, 1, 0, 1},
	// 	{0, 1, 0, 1, 0},
	// 	{0, 1, 1, 1, 0},
	// 	{1, 0, 0, 1, 1},
	// 	{0, 0, 1, 1, 0}}
	// rowCheck={}
	// colCheck=
	maxRow := len(c)
	maxCol := len(c[0])
	count := 0
	connected := 0
	// visit := make(map[[2]int]int)
	for i := 0; i <= maxRow-1; i++ {
		for j := 0; j <= maxCol-1; j++ {
			// fmt.Printf("Grid %+v\n", c)
			if c[i][j] == 1 {
				count = dfsCountv2(i, j, maxRow, maxCol, c)
				fmt.Printf("count %+v\n", count)
				if count >= 2 {
					connected++
				}
			}
		}
	}

	fmt.Printf("result %+v\n", connected)

}

func dfsCount(r int, c int, maxRow int, maxCol int, grid [][]int) int {
	connect := 0
	fmt.Printf("Traverse Grid[%d][%d]\n", r, c)
	// {0, 0, 1, 0, 0},
	// {0, 1, 0, 1, 0},
	// {0, 1, 0, 0, 0},
	// {1, 0, 0, 1, 1},
	// {0, 0, 1, 1, 0}}
	for i := 0; i < maxCol; i++ {
		if grid[r][i] == 1 {
			grid[r][i] = 0
			connect += 1
			connect += dfsCount(r, i, maxRow, maxCol, grid)
		}
	}

	for i := 0; i < maxRow; i++ {
		if grid[i][c] == 1 {
			grid[i][c] = 0
			connect += 1
			connect += dfsCount(i, c, maxRow, maxCol, grid)
		}
	}

	return connect
}

func dfsCountv2(r int, c int, maxRow int, maxCol int, grid [][]int) int {
	connect := 0
	fmt.Printf("Traverse Grid[%d][%d]\n", r, c)
	for i := 0; i < maxCol; i++ {
		if grid[r][i] == 1 {
			connect++
		}
	}
	if connect > 1 {
		return connect
	}
	connect = 0
	for i := 0; i < maxRow; i++ {
		if grid[i][c] == 1 {
			connect++
		}
	}

	return connect
}
