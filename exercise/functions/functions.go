//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func message() string {
	return "Hi there!"
}

func sum(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

func returnAnyNumber(number int) int {
	return number
}

func returnAnyTwoNumbers(number1, number2 int) (int, int) {
	return number1, number2
}

func main() {
	fmt.Println(greet("John Doe"))
	fmt.Println("Message:", message())
	fmt.Println(sum(5, 10, 16))
	fmt.Println(returnAnyNumber(14))
	fmt.Println(returnAnyTwoNumbers(1, 2))

	fmt.Println(returnAnyTwoNumbers(sum(1, 2, 3), 10))
}
