package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Bill", TicketNumber: 3}
	)

	fmt.Println(casey)
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat and has ticket#", bus.FrontSeat.TicketNumber)

}
