package main

import (
	"fmt"
)

func main() {
	var x, t, soma, k, par, maior, menor  int
	var e string

	for {
		fmt.Println("Digite um valor inteiro e positivo:")
		fmt.Scan(&x)
		t++         
		soma += x    

		if x%2 == 0 {
			par++     
			k += x    
		}

		if x > maior {
			maior = x 
		}
		if x < menor {
			menor = x 
		}

		fmt.Println("Deseja continuar? (s para sim, qualquer outra tecla para não):")
		fmt.Scan(&e)

		if e = 30000{
			break
		}
	}

	var media1, media2 float64

	if t > 0 {
		media1 = float64(soma) / float64(t)
	}
	if par > 0 {
		media2 = float64(k) / float64(par)
	}
	fmt.Println("\n--- Resultados ---")
	fmt.Println("Quantidade de valores digitados:", t)
	fmt.Println("Soma total:", soma)
	fmt.Printf("Média geral: %.2f\n", media1)
	fmt.Printf("Média dos pares: %.2f\n", media2)
	fmt.Println("Maior valor digitado:", maior)
	fmt.Println("Menor valor digitado:", menor)
}