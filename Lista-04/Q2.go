package main

import "fmt"

func main() {
	var v1 [10]int
	var v2 [5]int

	fmt.Println("Digite os 10 valores do primeiro vetor (v1):")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}

	fmt.Println("Digite os 5 valores do segundo vetor (v2):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&v2[i])
	}
	fmt.Println("O valor do primeiro vetor resultante:")
	fmt.Print("[")
	for i := 0; i < 10; i++ {

		if v1[i]%2 == 0 {

			s := v2[0] + v2[1] + v2[2] + v2[3] + v2[4]

			fmt.Print(" ", v1[i]+s, " ")

		}

	}
	fmt.Print("]")
	fmt.Println("\nO valor do segundo vetor resultante:")
	fmt.Print("[")
	for i := 0; i < 10; i++ {

		if v1[i]%2 != 0 {

			s := v2[0] + v2[1] + v2[2] + v2[3] + v2[4]

			fmt.Print(" ", v1[i]+s, " ")

		}

	}
	fmt.Print("]")

}
