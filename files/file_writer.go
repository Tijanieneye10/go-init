package files

import (
	"fmt"
	"os"
)

func WriteToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	err := os.WriteFile("balances.txt", []byte(balanceText), 0644)

	if err != nil {
		fmt.Println(err)
	}
}
