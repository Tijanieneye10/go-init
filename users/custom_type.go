package users

import "fmt"

type str string

func (s str) log() string {
	output := fmt.Sprintf("%s log here", s)
	return output
}

func main() {

	var name str = "John Doe"

	fmt.Println(name.log())

}
