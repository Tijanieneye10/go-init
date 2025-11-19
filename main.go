package main

import "learning/payments"

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

}

func InitializePayment(gateway Paymentable){
	gateway.Initialize()
}

func Pay(gateway Paymentable){
	gateway.Pay()
}