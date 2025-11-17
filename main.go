package main

import (
	"fmt"
	"learning/exceptions"
	"learning/files"
)

func parseInput(input string) (string, error) {
	if len(input) == 0 {
		return "", &exceptions.CustomError{Msg: "Empty input"}
	}

	return input, nil
}

func main() {

	files.WriteToFile(1110.00)
	result, err := parseInput("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
