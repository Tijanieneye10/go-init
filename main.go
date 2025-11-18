package main

import (
	"fmt"
	"learning/exceptions"
)

type MyError struct {
	Msg string
}

func (m *MyError) Error() string {
	return m.Msg
}

func parseInput(input string) (string, error) {
	if len(input) == 0 {
		return "", &exceptions.CustomError{Msg: "Empty input"}
	}

	return input, nil
}

func main() {

	fmt.Println("Hello World")
	result, _ := fmt.Scan("How are you doing?")

	fmt.Println(result)

}
