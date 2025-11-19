package payments

import "fmt"


type Flutterwave struct {
	Reference int
	Amount    float64
	Currency  string
}


func (p *Flutterwave) Initialize() {
	fmt.Println("Flutterwave initialized")
}


func (p *Flutterwave) Pay() {
	fmt.Println("Flutterwave paid")
}
