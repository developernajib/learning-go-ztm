//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	commandCount := 0
	lineCount := 0

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "Q" || input == "q" {
			break
		}

		if input != "" {
			lineCount++
		}

		switch input {
		case "hello":
			commandCount++
			fmt.Println("Hello there!")

		case "bye":
			commandCount++
			fmt.Println("Goodbye!")

		case "":
			continue

		default:
			commandCount++
			fmt.Println("Unknown command")
		}
	}

	fmt.Println("\nUsage Statistics")
	fmt.Println("----------------")
	fmt.Printf("Non-blank lines entered: %d\n", lineCount)
	fmt.Printf("Commands entered: %d\n", commandCount)
}
