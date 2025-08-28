//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"unicode"
)

func countLetters(word string) int {
	count := 0
	for _, r := range word {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var wg sync.WaitGroup
	results := make(chan int)

	for scanner.Scan() {
		word := scanner.Text()

		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			results <- countLetters(w)
		}(word)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0

	for v := range results {
		total += v
	}

	fmt.Println(total)
}
