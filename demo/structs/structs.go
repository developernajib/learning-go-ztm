package main

import "fmt"

type Bus struct {
	FrontSeat Passenger
}
type Passenger struct {
	Name         string
	TicketNumber uint
	Boarded      bool
}

func main() {
	// casey := Passenger{
	// 	Name:         "Casey",
	// 	TicketNumber: 1,
	// 	Boarded:      true,
	// }
	casey := Passenger{"Casey", 1, true}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	heidi.Boarded = true

	bill.Boarded = true

	if casey.Boarded {
		fmt.Println("Casey has boarded the bus")
	} else {
		fmt.Println("Casey has not boarded the bus")
	}
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	} else {
		fmt.Println("Bill has not boarded the bus")
	}
	if heidi.Boarded {
		fmt.Println("Heidi has boarded the bus")
	} else {
		fmt.Println("Heidi has not boarded the bus")
	}
	if ella.Boarded {
		fmt.Println("Ella has boarded the bus")
	} else {
		fmt.Println("Ella has not boarded the bus")
	}
	if !casey.Boarded && !bill.Boarded && !heidi.Boarded && !ella.Boarded {
		fmt.Println("No one has boarded the bus")
	}
}
