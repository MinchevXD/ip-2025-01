package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	soma := 0.00

	for i := 1.00; i <= x; i++ {
		soma += i
	}
	fmt.Println(soma)
}
