package main

import "fmt"

func main() {
	var v1 [10]int
	var menor, posicao int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&v1[i])
		if i == 0 || v1[i] < menor {
			menor = v1[i]
			posicao = i
		}
	}

	fmt.Println("O menor elemento do vetor é", menor, "e sua posição dentro do vetor é:", posicao)
}