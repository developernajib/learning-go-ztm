//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func report(lines []string) {
	letters := 0
	digits := 0
	spaces := 0
	punct := 0

	processRune := func(r rune) {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsPunct(r):
			punct++
		}
	}

	for _, line := range lines {
		for _, r := range line {
			processRune(r)
		}
	}

	fmt.Println("Report:----")
	fmt.Println("Letters:", letters)
	fmt.Println("Digits:", digits)
	fmt.Println("Spaces:", spaces)
	fmt.Println("Punctuation:", punct)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	report(lines)
}
