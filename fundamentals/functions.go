package main

// Function with parameters and return value
func sum(number1 int8, number2 int8) int8 {
	return number1 + number2
}

// In Go one function can return multiple values
func multipleReturn(number1, number2 int8) (int8, int8) {
	// Using := to declare the variables like inferring the type
	sum := number1 + number2
	sub := number1 - number2
	return sum, sub
}

func main() {
	// Call the function and pass the arguments
	// The result will be stored in the variable declared like inferring the type
	sum30 := sum(10, 20)
	println(sum30)

	// Call the function with multiple return values
	multipleReturnSum, multipleReturnSub := multipleReturn(10, 20)
	println(multipleReturnSum, multipleReturnSub)
	// We can also ignore the value we don't need
	_, multipleReturnSub2 := multipleReturn(10, 20)
	println(multipleReturnSub2)

	// In Go, we can declare a function as a variable
	var function = func() {
		println("Text of function variable")
	}
	function()
}
