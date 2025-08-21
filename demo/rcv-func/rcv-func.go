package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(spaceNum int, lot *ParkingLot) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 5)}
	fmt.Println("initial:", lot)
	lot.occupySpace(1)
	occupySpace(2, &lot)
	fmt.Println("After occupied:", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)
}
