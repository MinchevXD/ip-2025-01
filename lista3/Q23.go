package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite o número de valores da sequência:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O número deve ser positivo ")
		return
	}

	x := 1000.0
	soma := 0.0

	for i := 1; i <= n; i++ {
		k := x / float64(i)

		if i%2 == 0 {
			soma -= k
		} else {
			soma += k
		}

		x -= 3
	}

	fmt.Printf("Resultado final: %.2f\n", soma)
}
