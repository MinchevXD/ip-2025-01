package main

import "fmt"

func main() {
	var Lucromax, x, y float64
	ingressos := 130.0
	preco := 6.0
	for i := 1; i <= 10; i++ {
		lucro := ingressos*preco - 300
		ingressos += 30
		preco -= 0.6

		fmt.Printf("\n%.0f", lucro)
		if Lucromax < lucro {
			Lucromax = lucro
		}
		if x*y < ingressos*preco {
			x = ingressos
			y = preco
		}

	}
	fmt.Printf("\n%.0f", Lucromax)
	fmt.Printf("\n%.0f  %.1f", x, y)
}
