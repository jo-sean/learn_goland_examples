package main

import "fmt"

func main() {
	var sum int
	fmt.Println("Sum is", sum)
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Increment. Sum is", sum)
	}
	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

}
