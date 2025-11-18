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

type str string

func (s str) log() string {
	output := fmt.Sprintf("%s log here", s)
	return output
}

func main() {

	var name str = "John Doe"

	fmt.Println(name.log())

	person := users.New("John", "Doe", "johndoe@gmail.com")

	fmt.Println(person.FirstName())

	//person := users.NewAdmin("super-admin", "Doe")
	//
	//fmt.Println(person.Role)
}
