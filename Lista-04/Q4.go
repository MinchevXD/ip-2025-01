package main

import "fmt"

func main() {
	var A [10]int
	fmt.Println("Digite 10 números inteiros:")

	for i := 0; i < 10; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Println("\nElementos repetidos e suas quantidades:")

	encontrou := false
	for i := 0; i < 10; i++ {
		cont := 1
		if A[i] != -1 { 
			for j := i + 1; j < 10; j++ {
				if A[i] == A[j] {
					cont++
					A[j] = -1 
				}
			}
			if cont > 1 {
				fmt.Printf("Número %d aparece %d vezes\n", A[i], cont)
				encontrou = true
			}
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número está repetido.")
	}
}