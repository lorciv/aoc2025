package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var grid [][]string

	// Read input grid
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	// Print the grid
	// for i := 0; i < len(grid); i++ {
	// 	for j := 0; j < len(grid[i]); j++ {
	// 		fmt.Print(grid[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }

	// Walk the grid
	accessible := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "@" {
				continue
			}

			adjacentRolls := 0
			for _, di := range []int{-1, 0, 1} {
				for _, dj := range []int{-1, 0, 1} {
					// Skip cell itself
					if di == 0 && dj == 0 {
						continue
					}

					ai, aj := i+di, j+dj

					// Skip adjacent cells outside of board
					if ai < 0 || ai > len(grid)-1 || aj < 0 || aj > len(grid[i])-1 {
						continue
					}

					if grid[ai][aj] == "@" {
						adjacentRolls++
					}
				}
			}
			if adjacentRolls < 4 {
				accessible++
			}
		}
	}
	fmt.Println(accessible)
}
