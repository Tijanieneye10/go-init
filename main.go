package main

import (
	"fmt"
	"learning/payments"
)

type Paymentable interface {
	Initialize()
	Pay()
}

func main() {
	paystack := payments.Paystack{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&paystack)
	Pay(&paystack)

	flutterwave := payments.Flutterwave{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&flutterwave)
	Pay(&flutterwave)

	unkown(3)
}

func InitializePayment(gateway Paymentable){
	gateway.Initialize()
}

func Pay(gateway Paymentable){
	gateway.Pay()
}


func unkown(param interface{}) {
	aint, aIsInt := param.(int)
	if aIsInt {
		fmt.Println("a is an int", aint)
		return
	}
	bint, bIsInt := param.(int)
	if bIsInt {
		fmt.Println("b is an int", bint)
		return
	}
	fmt.Println("Unknown type")
}