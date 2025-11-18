package main

import (
	"fmt"
)

type User struct {
	Firstname string
	Lastname  string
	State     string
	Age       int
}

type UserError struct {
	Msg string
}

func (e *UserError) Error() string {
	return e.Msg
}

func newUser(firstname, lastname, state string, age int) (*User, error) {

	if firstname == "" || lastname == "" || state == "" {
		return nil, &UserError{Msg: "Invalid arguments"}
	}

	return &User{
		Firstname: firstname,
		Lastname:  lastname,
		State:     state,
		Age:       age,
	}, nil
}

func main() {
	person, err := newUser("John", "Doe", "Doe", 30)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person.Lastname)
}
