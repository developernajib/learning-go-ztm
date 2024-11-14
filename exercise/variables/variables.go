//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier

//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color string = "Green"
	fmt.Println("My favorite color is", color)

	birthYear, age := 2003, 21
	fmt.Println("My birth year is", birthYear, "and I am", age, "years old")

	var (
		firstInitial = "M"
		lastInitial  = "N"
	)
	fmt.Println("My initials are", firstInitial, lastInitial)

	var ageInDays int = age * 365
	fmt.Println("My age in days is", ageInDays)

	var height float32 = 5.9
	fmt.Println("My height is", height, "meters")

	var (
		skill1 = "Blockchain Development"
		skill2 = "Web Development"
		skill3 = "Cybersecurity"
		skill4 = "Programming"
	)

	fmt.Printf("Here are my skills: %s, %s, %s, %s and many more\n", skill1, skill2, skill3, skill4)

	skip, the, _ := "Data", "is", "normal"
	fmt.Println(skip, the)
}

// ------------ Go Variables ------------

// 1. Using var Keyword
// With Explicit Type and Initial Value:

// var name string = "John"
// var age int = 30
// With Type Inference (Omitting Type):

// var name = "John" // infers string type
// var age = 30      // infers int type
// Without Initial Value (Default Zero Value):

// var name string // default value is ""
// var age int     // default value is 0

// 2. Using Short Variable Declaration (:=)
// Only within functions, the := shorthand can be used to declare and initialize variables with type inference.

// name := "John"
// age := 30

// 3. Declaring Multiple Variables at Once
// Using var with Type Inference:

// var name, age, city = "John", 30, "New York"
// Using var with Explicit Types:
// var name, age string // both are strings, empty by default

// 4. Using const for Constants
// Constants are immutable once defined, declared using const instead of var.

// go
// Copy code
// const pi = 3.14159
// const country = "USA"
// 5. Using Block Assignment with var and const
// Block assignment groups multiple declarations, making code more organized.

// Using var Block:

// go
// Copy code
// var (
//     name = "John"
//     age  = 30
//     city = "New York"
// )
// Using const Block:

// const (
//     pi       = 3.14159
//     country  = "USA"
//     language = "Go"
// )

// 6. Pointer Variables
// Declaring pointers involves using the * operator.

// var name *string // declaring a pointer to a string
// And to initialize it:

// str := "John"
// namePtr := &str // `&` provides the memory address of `str`

// 7. Blank Identifier _
// If you want to ignore a variable, Go provides the blank identifier _. This is commonly used for ignoring return values.

// _, err := someFunctionThatReturnsTwoValues()

// 8. Using new to Allocate Memory
// The new function allocates memory and returns a pointer to the zero value of the specified type.

// namePtr := new(string)
// *namePtr = "John"

// 9. Using make for Slices, Maps, and Channels
// The make function initializes slices, maps, and channels with allocated memory, returning the initialized object.

// mySlice := make([]int, 5)      // Slice of integers, length 5
// myMap := make(map[string]int)  // Map with string keys, int values
// myChan := make(chan int)       // Channel of integers

// 10. Using Function Return Values as Variables
// Go allows direct assignment of function return values to variables.

// name, age := getNameAndAge()
