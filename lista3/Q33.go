package main

import "fmt"

func main() {
	var n1, n2 int
	soma := 0
	k := 0.0
	fmt.Scan(&n2, &n1)
	for i := 0; i <= n2; i += n1 {
		soma = i
		k++

	}
	fmt.Println(n2-soma, k-1)
}
