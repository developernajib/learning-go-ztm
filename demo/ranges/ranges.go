package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	// for _, element := range slice {
	// 	fmt.Println(element)
	// }

	for i, element := range slice {
		fmt.Println(i, element, ":")

		for _, char := range element {
			fmt.Printf(" %q", char)
		}
		fmt.Printf("\n\n")
	}
}
