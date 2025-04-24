package main

import (
	"fmt"
)

func main() {
	var H float64 = 0
	num := 1

	for den := 1; den <= 50; den++ {
		H += float64(num) / float64(den)
		num += 2
	}

	fmt.Println(H)
}
