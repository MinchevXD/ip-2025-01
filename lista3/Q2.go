package main

import "fmt"

func main() {
	for i := 0; i <= 20; i += 2 {
		soma := 50 + i
		f := 0
		f += soma
		media := f / 10
		fmt.Println(f, media)
	}
}
