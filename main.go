package main

import "fmt"

type Product struct {
	ID          int
	name        string
	description string
}

func main() {

	names := make(map[string]int)

	fmt.Println(len(names))

	products := []Product{Product{
		ID:          1,
		name:        "Product 1",
		description: "Description 1",
	},
		Product{
			ID:          2,
			name:        "Product 2",
			description: "Description 2",
		},
	}

	fmt.Println(products[0].name)

	userDetails := map[string]string{
		"name": "John Doe",
		"age":  "30",
		"city": "New York",
	}

	fmt.Println(userDetails["name"])
	delete(userDetails, "age")
	fmt.Println(userDetails["age"])

}
