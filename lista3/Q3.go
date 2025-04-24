package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&a)
	b := a / 3
	c := a
	d := b

	e := 0
	for d < c {
		c *= 1.02
		d *= 1.05
		e++
	}
	fmt.Printf("\nO salário derá ultrapassado  em %d meses.\n", e)
}
