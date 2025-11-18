package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	State     string
	Age       int
}

func newUser(firstname, lastname, state string, age int) *User {
	return &User{
		Firstname: firstname,
		Lastname:  lastname,
		State:     state,
		Age:       age,
	}
}

func main() {
	person := newUser("John", "Doe", "Doe", 30)

	fmt.Println(person.Lastname)
}
