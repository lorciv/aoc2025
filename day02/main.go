package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	line := string(buf)

	ranges := strings.Split(line, ",")

	sum := 0
	for _, r := range ranges {
		var first, last int
		fmt.Sscanf(r, "%d-%d", &first, &last)

		for id := first; id <= last; id++ {
			idStr := strconv.Itoa(id)

			if len(idStr)%2 == 1 {
				continue
			}
			middle := len(idStr) / 2

			left, right := idStr[:middle], idStr[middle:]
			if left == right {
				sum += id
			}
		}
	}

	fmt.Println(sum)
}
