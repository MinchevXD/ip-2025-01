package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	for i := 1.00; i <= float64(n1*n2); i++ {

		if int(i)%n1 == 0 && int(i)%n2 == 0 {
			fmt.Println(i)
			break
		}
	}
}
