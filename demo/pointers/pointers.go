package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter:", counter.hits)
}

func replace(counter *Counter, old *string, new string) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&counter, &world, "There!")
	fmt.Println(hello, world)

	// Updating the phrase variable not the old ones
	phrase := []string{hello, world}
	replace(&counter, &phrase[0], "Hi")
	fmt.Println(phrase)
}
