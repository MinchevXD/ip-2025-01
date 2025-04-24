package main

import (
	"fmt"
	"math"
)

func main() {

	soma := 0.00
	for i := 0; i <= 51; i++ {
		k := (2*i + 1) * (2*i + 1) * (2*i + 1)
		if i%2 == 0 {
			soma += 1.0 / float64(k)
		} else {
			soma -= 1.0 / float64(k)
		}

	}
	y := soma * 32
	c := math.Cbrt(y)
	fmt.Println(c)
}
