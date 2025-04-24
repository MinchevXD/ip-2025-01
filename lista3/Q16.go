package main

import "fmt"

func main() {
	var n1, n2, k int
	fmt.Println("Digite dois números naturais:")
	fmt.Scan(&n1, &n2)
	fmt.Println("Digite a quantidade de termos que deseja gerar:")
	fmt.Scan(&k)
	var v = make([]float64, k)
	v[0] = float64(n1)
	v[1] = float64(n2)
	for i := 2; i < k; i++ {
		if i%2 == 0 {
			v[i] = v[i-1] + v[i-2]
		} else {
			v[i] = v[i-1] - v[i-2]
		}
	}
	fmt.Println("Sequência gerada:")
	for i := 0; i < k; i++ {
		fmt.Printf("%.0f ", v[i])
	}
	fmt.Println()
}
