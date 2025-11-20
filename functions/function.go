package functions



func TransformNumber(number int, transform func(int) int) int {
	return transform(number)
}



func Double(number int) int {
	return number * 2
}

func Triple(number int) int {
	return number * 3
}