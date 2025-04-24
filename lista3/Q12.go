package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Print("Digite o valor de X: ")
	fmt.Scan(&x)

	var soma float64 = 0
	sinal := 1.0
	fatorial := 1.0

	for n := 0; n < 20; n++ {
		if n > 0 {
			fatorial *= float64(n)
		}
		termo := sinal * x / fatorial
		soma += termo
		sinal *= -1
	}

	fmt.Printf("Resultado do somatório: %.10f\n", soma)
}
