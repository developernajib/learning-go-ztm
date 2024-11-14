package main

import "fmt"

func add(number1, number2 int) int {
	return number1 + number2
}

func greet(name string) string {
	return "Hello, " + name
}

func main() {
	result := add(2, 9)
	fmt.Println("The sum is", result)

	fmt.Println(greet("John Doe"))
}
