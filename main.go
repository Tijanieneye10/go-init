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
	person := users.NewAdmin("John", "Doe")

	fmt.Println(person.FirstName())
}
