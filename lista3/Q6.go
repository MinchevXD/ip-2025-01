package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("Digite um valor inteiro e positivo:")
	fmt.Scan(&x)

	var y int

	for i := 1; i*(i+1)*(i+2) <= x; i++ {
		if x == i*(i+1)*(i+2) {
			y = i
			break
		}
	}

	if y > 0 {
		fmt.Printf("O número é triângular")
	} else {
		fmt.Println("O número não é triângular.")
	}
}
