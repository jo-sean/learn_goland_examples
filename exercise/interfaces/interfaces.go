//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//

//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

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

type Vehicle string

const (
	small  Vehicle = "motorcycles"
	medium Vehicle = "cars"
	large  Vehicle = "trucks"
)

type Vehicles struct {
	carType Vehicle
	name    string
}

func (v *Vehicles) Lift() {
	switch v.carType {
	case small:
		fmt.Printf("%s is a %s. The vehicle was moved to the small lift.\n", v.name, v.carType)
	case medium:
		fmt.Printf("%s is a %s. The vehicle was moved to the medium lift.\n", v.name, v.carType)
	case large:
		fmt.Printf("%s is a %s. The vehicle was moved to the large lift.\n", v.name, v.carType)
	default:
		fmt.Println("Vehicle type not supported")
	}

}

type Servicer interface {
	Lift()
}

func main() {
	moto := Vehicles{small, "Harley Davidson"}
	car := Vehicles{medium, "Toyota Rav4"}
	truck := Vehicles{large, "Ford F-150"}

	shop := []Servicer{&moto, &car, &truck}

	// 5. Loop through and execute the interface method
	for _, vehicle := range shop {
		vehicle.Lift()
	}

}
