package main

import "fmt"

func main() {
	var v [201]int
	for i := 1; i <= 200; i += 2 {
		v[i] = i
		fmt.Println(v[i])
	}
}
