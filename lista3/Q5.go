package main

import (
	"fmt"
)

func main() {
	var x, y, z float64
	var e string

	var total int
	var v int
	var m int
	var n int
	var soma float64

	for {
		fmt.Println("Digite o valor da altura, idade e peso respectivamente:")
		fmt.Scan(&x, &y, &z)

		total++

		if y > 50 {
			v++
		}

		if y >= 10 && y <= 20 {
			soma += x
			m++
		}

		if z < 40 {
			n++
		}

		fmt.Println("Digite 1 para continuar colocando dados e qualquer outro valor para parar:")
		fmt.Scan(&e)

		if e != "1" {
			break
		}
	}

	fmt.Printf("\nQuantidade de pessoas com mais de 50 anos: %d\n", v)

	if m > 0 {
		media := soma / float64(m)
		fmt.Printf("Média das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", media)
	} else {
		fmt.Println("Nenhuma pessoa com idade entre 10 e 20 anos foi informada.")
	}

	if total > 0 {
		porcentagem := (float64(n) / float64(total)) * 100
		fmt.Printf("Porcentagem de pessoas com peso inferior a 40kg: %.2f%%\n", porcentagem)
	}
}
