package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tile struct {
	x, y int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func area(t, u tile) int {
	dx := abs(t.x-u.x) + 1
	dy := abs(t.y-u.y) + 1
	return dx * dy
}

func main() {
	var tiles []tile
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			log.Fatalf("could not parse input line: %v", err)
		}
		tiles = append(tiles, tile{x, y})
	}

	max := 0
	for i, t := range tiles {
		for _, u := range tiles[i+1:] {
			if a := area(t, u); a > max {
				max = a
			}
		}
	}
	fmt.Println(max)
}
