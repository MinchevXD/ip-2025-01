package main

import (
	"fmt"
)

func main() {
	var id, maisGordoID, maisMagroID int
	var peso, maisGordoPeso, maisMagroPeso float64

	maisGordoPeso = -1
	maisMagroPeso = 1e9

	for i := 1; i <= 90; i++ {
		fmt.Printf("Digite o número de identificação do boi #%d: ", i)
		fmt.Scan(&id)

		fmt.Printf("Digite o peso do boi #%d: ", i)
		fmt.Scan(&peso)

		if peso > maisGordoPeso {
			maisGordoPeso = peso
			maisGordoID = id
		}

		if peso < maisMagroPeso {
			maisMagroPeso = peso
			maisMagroID = id
		}
	}

	fmt.Println("\n--- Resultados ---")
	fmt.Printf("Boi mais gordo: ID %d com %.2f kg\n", maisGordoID, maisGordoPeso)
	fmt.Printf("Boi mais magro: ID %d com %.2f kg\n", maisMagroID, maisMagroPeso)

}
