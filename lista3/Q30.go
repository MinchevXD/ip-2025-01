package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0.00; i <= 20; i += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * math.Pow(i, 3)
		fmt.Printf("\nO volume da esfera de raio %.2f é %.2f", i, volume)
	}
}
