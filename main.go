package main

import (
	"fmt"
	"learning/functions"
	"learning/payments"
)

type Paymentable interface {
	Initialize()
	Pay()
}

// func doubleNumber(numbers *[]int) []int {
// 	dNumber := []int{}

// 	for _, num := range *numbers {
// 		dNumber = append(dNumber, num*2)
// 	}

// 	return dNumber
// }


func main() {

	result := functions.TransformNumber(4, functions.Triple)

	fmt.Println("here is the result:", result)


	// numbers := []int{1, 2, 3, 4, 5}

	// result := doubleNumber(&numbers)

	// fmt.Println(result)



	paystack := payments.Paystack{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&paystack)
	Pay(&paystack)

	flutterwave := payments.Flutterwave{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&flutterwave)
	Pay(&flutterwave)

	unkown(3)
	unkown("number")
}

func InitializePayment(gateway Paymentable){
	gateway.Initialize()
}

func Pay(gateway Paymentable){
	gateway.Pay()
}


func unkown(param any) {
	aint, aIsInt := param.(int)
	if aIsInt {
		fmt.Println("a is an int", aint)
		return
	}
	bstring, bIsString := param.(string)
	if bIsString {
		fmt.Println("b is an string", bstring)
		return
	}
	fmt.Println("Unknown type")
}

func generic[T any](param T) T {
	fmt.Println(param)
	return param
}
