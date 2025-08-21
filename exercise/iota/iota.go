//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Addition = iota
	Subtraction
	Multiplication
	Division
)

type Operation int

func (op Operation) calculate(first, second int) int {
	switch op {
	case Addition:
		return first + second
	case Subtraction:
		return first - second
	case Multiplication:
		return first * second
	case Division:
		return first / second
	}
	panic("Unknown operation")
}

func main() {
	add := Operation(Addition)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operation(Subtraction)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Operation(Multiplication)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operation(Division)
	fmt.Println(div.calculate(100, 2)) // = 50
}
