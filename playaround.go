//package main
//
//import (
//	"fmt"
//	"learning/exceptions"
//)
//
//type MyError struct {
//	Msg string
//}
//
//type Name string
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (n Name) MyName() string {
//	return string(n)
//}
//
//func (p *Person) GetDetail() string {
//	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
//}
//
//func (m *MyError) Error() string {
//	return m.Msg
//}
//
//func parseInput(input string) (string, error) {
//	if len(input) == 0 {
//		return "", &exceptions.CustomError{Msg: "Empty input"}
//	}
//
//	return input, nil
//}
//
//func main() {
//
//	fmt.Println(Name.MyName("Tijani"))
//
//	person := Person{Name: "Usman", Age: 20}
//
//	fmt.Println(person.GetDetail())
//
//	age := 18
//	agePointer := &age
//
//	fmt.Println(getAge(agePointer))
//
//	fmt.Println("Hello World")
//	result, _ := fmt.Scan("How are you doing?")
//
//	fmt.Println(result)
//
//}
//
//func getAge(age *int) int {
//	return *age - 3
//}
