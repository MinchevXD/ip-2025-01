package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Println("Digite dois números naturais:")
	fmt.Scan(&n1, &n2)
	menor := n1
	maior := n2
	if n1 > n2 {
		menor = n2
		maior = n1
	}
	for k := menor + 1; k < maior; k++ {
		divisivel := 0
		for i := 2; i*i <= k; i++ {
			if k%i == 0 {
				divisivel = 1
				break
			}
		}

		if k > 1 && divisivel == 0 {
			fmt.Println(k)
		}
	}
}
