package main

import "fmt"

func main() {
	var x float64
	for {
		fmt.Scan(&x)
		if x <= 0 {
			break
		}
		quadrado := false
		for i := 1.0; i*i <= x; i++ {
			if i*i == x {
				quadrado = true
				break
			}
		}
		if quadrado {
			fmt.Println("O número é um quadrado perfeito.")
		} else {
			fmt.Println("O número NÃO é um quadrado perfeito.")
		}
	}
}
