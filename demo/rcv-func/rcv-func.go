package main

import "fmt"

type Space struct {
	occupied bool
}

type parkingLot struct {
	spaces []Space
}

func occupySpace(lot *parkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *parkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *parkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := parkingLot{spaces: make([]Space, 5)}
	fmt.Println(lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied: ", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate: ", lot)

}
