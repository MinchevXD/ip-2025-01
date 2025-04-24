package main

import (
	"fmt"
	"math"
)

func maclaurinSin(A float64) float64 {
	A3 := math.Pow(A, 3)
	A5 := math.Pow(A, 5)
	A7 := math.Pow(A, 7)

	return A - (A3 / 6.0) + (A5 / 120.0) - (A7 / 5040.0)
}

func main() {
	fmt.Println("Ângulo (rad)\tSérie Maclaurin\tmath.Sin\tDiferença")
	fmt.Println("-------------------------------------------------------------")

	for A := 0.0; A <= 6.3; A += 0.1 {
		A = math.Round(A*10) / 10
		maclaurin := maclaurinSin(A)
		realSin := math.Sin(A)
		difference := math.Abs(maclaurin - realSin)

		fmt.Printf("%.1f\t\t%.6f\t\t%.6f\t\t%.6f\n",
			A, maclaurin, realSin, difference)
	}
}
