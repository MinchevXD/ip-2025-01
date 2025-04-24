package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	hexChars := "0123456789abcdef"
	result := ""

	if x == 0 {
		result = "0"
	} else {
		for x > 0 {
			resto := x % 16
			result = string(hexChars[resto]) + result
			x /= 16
		}
	}

	fmt.Println(result)
}
