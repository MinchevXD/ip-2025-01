package main

import "fmt"

func main() {
	soma := 0.00
	k := 1.00
	total := 0.00
	for i := 1; i <= 20; i++ {
		soma += 100 - float64(i)
		k *= float64(i)
		termo := soma / k
		total += termo

	}
	fmt.Println(total)
}
