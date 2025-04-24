package main

import (
	"fmt"
)

func main() {
	s := 0.0
	numerador := 1.0

	for i := 0; i < 15; i++ {
		denominador := (15 - i) * (15 - i)
		s += numerador / float64(denominador)
		numerador *= -2.0
	}

	fmt.Printf("O valor de S é: %.15f\n", s)
}
