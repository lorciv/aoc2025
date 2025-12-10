package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type interval struct {
	low, high int
}

func main() {
	var intervals []interval

	// Read ID intervals of fresh items
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		var l, h int
		fmt.Sscanf(line, "%d-%d", &l, &h)
		intervals = append(intervals, interval{low: l, high: h})
	}

	// Count IDs of fresh items
	count := 0
	for scan.Scan() {
		line := scan.Text()
		id, _ := strconv.Atoi(line)
		for _, t := range intervals {
			if id >= t.low && id <= t.high {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}
