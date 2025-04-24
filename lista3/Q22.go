package main

import "fmt"

func main() {
	var soma int

	for y := 1; y <= 37; y++ {
		x := (39 - y) * (38 - y)
		soma += x / y
	}

	fmt.Println(soma)
}
