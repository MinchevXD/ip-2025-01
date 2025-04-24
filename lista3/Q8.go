package main

import (
	"fmt"
)

func main() {
	var soma int

	fmt.Println("Números de 1 a 20:")

	for i := 1; i <= 20; i++ {
		fmt.Println(i)
		soma += i
	}

	fmt.Println("Soma dos números de 1 a 20:", soma)
}
