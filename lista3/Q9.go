package main

import (
	"fmt"
)

func main() {
	var N int
	var x, y float64
	var m1, m2, m3 int
	var k float64
	fmt.Print("Digite o número de alunos: ")
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Digite as duas notas do aluno %d: ", i)
		fmt.Scan(&x, &y)
		media := (x + y) / 2
		k += media
		if media < 3 {
			fmt.Println(media, "Reprovado")
			m1++
		} else if media >= 3 && media < 7 {
			fmt.Println(media, "Exame")
			m2++
		} else {
			fmt.Println(media, "Aprovado")
			m3++
		}
	}
	mediaC := k / float64(N)
	fmt.Println("Total de alunos aprovados:", m3)
	fmt.Println("Total de alunos de exame:", m2)
	fmt.Println("Total de alunos reprovados:", m1)
	fmt.Printf("Média da classe: %.2f\n", mediaC)
}
