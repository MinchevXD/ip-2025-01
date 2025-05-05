package main
import "fmt"
func main() {
	var v1 [101]int
	for i := 100; i >=0; i-- {
	v1[i]=i
	fmt.Println(v1[i])
	}
}
