package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	soma := 1.0
	for i := 0.0; i < y; i++ {
		soma *= x
	}

	fmt.Printf("%.0f\n", soma)
}
