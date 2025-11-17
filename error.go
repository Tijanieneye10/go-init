package main

import (
	"errors"
	"fmt"
)

func getPositiveNumber(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("invalid number provided")
	}

	return num, nil
}

func run() {
	num, err := getPositiveNumber(21)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
