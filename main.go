package main

import (
	"fmt"
	"learning/users"
)

type UserError struct {
	Msg string
}

func (e *UserError) Error() string {
	return e.Msg
}

func main() {
	person := users.New("John", "Doe", "johndoe@gmail.com")

	fmt.Println(person.FirstName())
}
