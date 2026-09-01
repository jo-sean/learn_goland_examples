package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap item!")
	case p < 10:
		fmt.Println("Moderately cheap item")\
	default:
		fmt.Println("Expensive item")
	}


	switch ticket := Economy; {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First Class seating")
	default:
		fmt.Println("Other seating")
	}

}
