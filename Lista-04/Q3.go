package main 
import "fmt"
func main() {
	var v1 [10]int
	soma1:=0
	soma2:=0
	fmt.Println("Digite 10 valores inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}
	fmt.Print("a)- [")
	for i := 0; i < 10; i++ {
		if v1[i]%2 == 0{
			fmt.Print(" ", v1[i], " ")
		}
	}
	fmt.Print("]")
	fmt.Print("\nb)-")
	for i := 0; i < 10; i++ {
		if v1[i]%2 == 0{
			
			soma1+=v1[i]
			
		}
	}
	fmt.Print(" ", soma1, " ")
	fmt.Print("\nc)- [")
	for i := 0; i < 10; i++ {
		if v1[i]%2 != 0{
			fmt.Print(" ", v1[i], " ")
		}
	}
	fmt.Print("]")
	fmt.Print("\nd)-")
	for i := 0; i < 10; i++ {
		if v1[i]%2 != 0{
			
			soma2+=v1[i]
			
		}
	}
	fmt.Print(" ", soma2, " ")
}