package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite o número de termos da sequência:")
	fmt.Scan(&n)

	fmt.Println("Série dos quadrados perfeitos até o", n, "º termo:")
	for i := 1; i <= n; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}
