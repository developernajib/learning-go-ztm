package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	fmt.Println("Performing an inspection...")
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Kitchen"},
		{name: "Bedroom", cleaned: true},
	}

	checkCleanliness(rooms)

	fmt.Println("\nCleaning rooms...")
	rooms[0].cleaned = true
	rooms[1].cleaned = true
	fmt.Println("Rooms cleaned")
	checkCleanliness(rooms)
}
