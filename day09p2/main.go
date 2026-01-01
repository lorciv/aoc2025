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

func print(floor map[tile]string) {
	var max tile
	for t := range floor {
		if t.x > max.x {
			max.x = t.x
		}
		if t.y > max.y {
			max.y = t.y
		}
	}

	for y := range max.y + 2 {
		for x := range max.x + 2 {
			t := floor[tile{x, y}]
			if t == "" {
				t = "."
			}
			fmt.Printf("%s ", t)
		}
		fmt.Print("\n")
	}
}

func main() {
	var redTiles []tile
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			log.Fatalf("could not parse input line: %v", err)
		}
		redTiles = append(redTiles, tile{x, y})
	}

	floor := make(map[tile]string)
	for _, t := range redTiles {
		floor[t] = "#"
	}

	for i := 0; i < len(redTiles); i++ {
		j := i + 1
		if j == len(redTiles) {
			j = 0
		}
		t, u := redTiles[i], redTiles[j]

		// Horizontal connection
		if t.y == u.y {
			if t.x > u.x {
				t, u = u, t
			}
			for k := t.x + 1; k < u.x; k++ {
				floor[tile{k, t.y}] = "X"
			}
		}

		// Vertical connection
		if t.x == u.x {
			if t.y > u.y {
				t, u = u, t
			}
			for k := t.y + 1; k < u.y; k++ {
				floor[tile{t.x, k}] = "X"
			}
		}
	}
	print(floor)

	// max := 0
	// for i, t := range tiles {
	// 	for _, u := range tiles[i+1:] {
	// 		if a := area(t, u); a > max {
	// 			max = a
	// 		}
	// 	}
	// }
	// fmt.Println(max)
}
