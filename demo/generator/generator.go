package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(r *rand.Rand, min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- r.Intn(max-min+1) + min
		}
	}()
	return out
}

func generateRandIntn(r *rand.Rand, count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- r.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randInt := generateRandInt(r, 1, 100)

	fmt.Println("generateRandInt infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("generateRandInt using finite count")
	randIntnRange := generateRandIntn(r, 2, 1, 10)

	for i := range randIntnRange {
		fmt.Println(i)
	}

	fmt.Println("generateRandInt using finite loop")
	randIntn := generateRandIntn(r, 3, 1, 10)
	for {
		n, open := <-randIntn
		if !open {
			break
		}
		fmt.Println(n)
	}
}
