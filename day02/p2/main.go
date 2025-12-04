package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// An ID is invalid if it is made only of some sequence of digits repeated at least twice.
// So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times),
// and 1111111 (1 seven times) are all invalid IDs.
func valid(id int) bool {
	idStr := strconv.Itoa(id)
	for winlen := 1; winlen < len(idStr)/2; winlen++ {
		if len(idStr)%winlen != 0 {
			continue
		}

		repeated := true
		for winstart := 0; winstart < len(idStr)-winlen; winstart += winlen {
			a := idStr[winstart : winstart+winlen]
			b := idStr[winstart+winlen : winstart+(2*winlen)]
			if a != b {
				repeated = false
				break
			}
		}
		if repeated {
			return false
		}
	}
	return true
}

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
			fmt.Println(id, valid(id))
		}
	}

	fmt.Println(sum)
}
