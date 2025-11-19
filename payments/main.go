// package main

// import (
// 	"fmt"
// 	"learning/payments"
// )

// type Paymentable interface {
// 	Initialize()
// 	Pay()
// }

// func main() {
// 	paystack := payments.Paystack{Reference: 1, Amount: 100, Currency: "NGN"}
// 	InitializePayment(&paystack)
// 	Pay(&paystack)

// 	flutterwave := payments.Flutterwave{Reference: 1, Amount: 100, Currency: "NGN"}
// 	InitializePayment(&flutterwave)
// 	Pay(&flutterwave)

// 	unkown(3)
// 	unkown("number")
// }

// func InitializePayment(gateway Paymentable){
// 	gateway.Initialize()
// }

// func Pay(gateway Paymentable){
// 	gateway.Pay()
// }


// func unkown(param any) {
// 	aint, aIsInt := param.(int)
// 	if aIsInt {
// 		fmt.Println("a is an int", aint)
// 		return
// 	}
// 	bstring, bIsString := param.(string)
// 	if bIsString {
// 		fmt.Println("b is an string", bstring)
// 		return
// 	}
// 	fmt.Println("Unknown type")
// }

// func generic[T any](param T) T {
// 	fmt.Println(param)
// 	return param
// }