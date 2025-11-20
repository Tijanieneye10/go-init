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

	annonymous := functions.TransformNumber(5, func(number int) int {
		return number * 4
	})

	fmt.Println("here is annonymous fn:", annonymous)

	myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sumNumbers(myNumbers...))



	paystack := payments.Paystack{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&paystack)
	Pay(&paystack)

	flutterwave := payments.Flutterwave{Reference: 1, Amount: 100, Currency: "NGN"}
	InitializePayment(&flutterwave)
	Pay(&flutterwave)

	unkown(3)
	unkown("number")
}


func sumNumbers(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
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
