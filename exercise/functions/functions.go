//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import (
	"fmt"
	random "math/rand/v2"
)

// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func namedGreeting(name string) {
	fmt.Println("Welcome to this lil program", name+"!")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func returnMessage() string {
	return "Schweet!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(numbers ...float64) float64 {
	var answer float64
	for _, value := range numbers {
		answer += value
	}
	return answer
}

// * Write a function that returns any number
func anyNumber() float64 {
	return random.Float64() * 100
}

// * Write a function that returns any two numbers
func anyTwoNumbers() (float64, float64) {
	return random.Float64() * 100, random.Float64() * 100
}

func main() {

	var num1, num2, num3 float64
	var name string

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	num1 = anyNumber()
	num2, num3 = anyTwoNumbers()

	fmt.Println("\n#################################################################")
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	namedGreeting(name)
	fmt.Printf("The three random numbers are %f, %f, and %f.\n", num1, num2, num3)
	fmt.Println("The sum of these numbers is", add(num1, num2, num3))
	fmt.Println("#################################################################\n")

	//* Call every function at least once

}
