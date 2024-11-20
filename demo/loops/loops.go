package main

import "fmt"

func main() {
	sum := 0
	// fmt.Println("Sum is", sum)
	for i := 0; i <= 100; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum >= 500 {
		sum -= 10
		fmt.Println("Decrement Sum is", sum)
	}
}
