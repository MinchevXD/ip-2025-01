package main

import (
	"fmt"
)

func main() {
	var b, n int
	fmt.Println("Digite os valores da base e do expoente ")
	fmt.Scan(&b, &n)
	total := b
	for i := 1; i < n; i++ {

		total *= b
	}

	fmt.Println(total)

}
