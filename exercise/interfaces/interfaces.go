//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vehicle interface {
	GetType() string
}

type Motorcycle struct{}

func (m Motorcycle) GetType() string {
	return "Motorcycle"
}

type Car struct{}

func (c Car) GetType() string {
	return "Car"
}

type Truck struct{}

func (t Truck) GetType() string {
	return "Truck"
}

func directToLift(v Vehicle) {
	var lift string

	switch v.GetType() {
	case "Motorcycle":
		lift = "Small Lift"
	case "Car":
		lift = "Standard Lift"
	case "Truck":
		lift = "Large Lift"
	default:
		lift = "Unknown Lift"
	}

	fmt.Printf("Vehicle Type: %s\n", v.GetType())
	fmt.Printf("Assigned Lift: %s\n\n", lift)
}

func main() {
	motorcycle := Motorcycle{}
	car := Car{}
	truck := Truck{}

	directToLift(motorcycle)
	directToLift(car)
	directToLift(truck)
}
