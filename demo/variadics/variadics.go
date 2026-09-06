package main

import "fmt"

func sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	// Must add the ... to the second argument in append() to add all the remainig numbers in the slice.
	// If not it'll be only the first one
	all := append(a, b...)

	// Must add the ... to the second argument in a variadic finction if it is a slice
	// to add all the remainig numbers in the slice. If not it'll through a runtime error.
	// The three dots (...) essentially does this sum(1, 2, 3, 4, 5, 6)
	answer := sum(all...)

	fmt.Println(answer)

}
