package main

import "fmt"

func main() {
	var N int
	fmt.Print("Digite um valor: ")
	fmt.Scan(&N)

	f := 1
	for i := 1; i <= N; i++ {
		f *= i
	}
	fmt.Println(f)
	if N < 0 {
		fmt.Println("Valor invalido")
	}
}
