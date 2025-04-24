package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, i int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i)

	soma := i*2 + h*3 + g*4 + f*5 + e*6 + d*7 + c*8 + b*9 + a*10
	y := soma % 11
	if y <= 2 {
		y = 0
	} else {
		y = 11 - y
	}

	soma2 := y*2 + i*3 + h*4 + g*5 + f*6 + e*7 + d*8 + c*9 + b*10 + a*11
	y2 := soma2 % 11
	if y2 <= 2 {
		y2 = 0
	} else {
		y2 = 11 - y2
	}

	fmt.Printf("%d%d%d.%d%d%d.%d%d%d-%d%d\n", a, b, c, d, e, f, g, h, i, y, y2)
}
