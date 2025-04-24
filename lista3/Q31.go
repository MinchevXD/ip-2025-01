package main

import "fmt"

func main() {
	k := 0.0
	y := 1.0

	for i := 1.0; i < 64; i++ {
		y *= 2
		k += y
		fmt.Println(i, y, k)
	}
	fmt.Println(k + 1)
}
