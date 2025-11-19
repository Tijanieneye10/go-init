package payments

import "fmt"


type Paystack struct {
	Reference int
	Amount    float64
	Currency  string
}


func (p *Paystack) Initialize() {
	fmt.Println("Paystack initialized")
}


func (p *Paystack) Pay() {
	fmt.Println("Paystack paid")
}
