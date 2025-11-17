package files

import (
	"fmt"
	"os"
	"strconv"
)

func WriteToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	err := os.WriteFile("balances.txt", []byte(balanceText), 0644)

	if err != nil {
		fmt.Println(err)
	}
}

func ReadFromFile() float64 {
	file, err := os.ReadFile("balances.txt")

	if err != nil {
		fmt.Println(err)
	}

	stringBalance := string(file)

	balance, _ := strconv.ParseFloat(stringBalance, 64)

	return balance
}
