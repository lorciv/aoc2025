package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read grid
	var grid [][]string
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	splitters := 0
	for i := range grid {
		// Skip first row
		if i == 0 {
			continue
		}

		// Propagate beams
		for j := range grid[i] {
			if grid[i-1][j] == "S" || grid[i-1][j] == "|" {
				switch grid[i][j] {
				case ".":
					grid[i][j] = "|"
				case "^":
					grid[i][j-1] = "|"
					grid[i][j+1] = "|"
					splitters++
				}
			}
		}
	}

	// Print the grid
	// for i := range grid {
	// 	for j := range grid[i] {
	// 		fmt.Print(grid[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }

	fmt.Println(splitters)
}
