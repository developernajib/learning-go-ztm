package main

import "fmt"

func main() {
	var x = 5                     // variable can be set without type (not recommanded)
	fmt.Println("Value of x:", x) // the command will set an speace after the x: automatically

	var fullName string = "Md. Najib Islam"
	fmt.Println("My name is", fullName)

	role := "Admin" // create and assign
	fmt.Println("My role is", role)

	var number int // unassinged variable (value: 0, because of int data type)
	fmt.Println("The sum is", number)

	part1, other := 1, 4
	fmt.Println("Part1 is", part1, "Other is", other)

	part2, other, partNone := 5, 0, 9
	fmt.Println("Part2 is", part2, "Other is", other, "PartNone is", partNone)

	part3, part1 := 10, 16                                 // reassign will work fine
	fmt.Printf("Part3 is %d, Part1 is %d\n", part3, part1) // C format so that we can use the comma on the desire place

	sum := part1 + part3
	fmt.Println("The sum is", sum)

	var ( // Block assignment
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson name is", lessonName, "and type is", lessonType)

	word1, word2, _ := "Hello", "World", "!" // compound assignment for ignoring the value
	fmt.Println(word1, word2)
}
