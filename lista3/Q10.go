package main

import "fmt"

func main() {
	var N int
	fmt.Println("Digite o número de termos da sequência de Fibonacci (N >= 3):")
	fmt.Scan(&N)
	vetor := make([]int, N)
	vetor[0] = 0
	vetor[1] = 1
	fmt.Println(vetor[0])
	fmt.Println(vetor[1])
	for i := 2; i < N; i++ {
		vetor[i] = vetor[i-2] + vetor[i-1]
		fmt.Println(vetor[i])
	}
}
