//--Summary:
//  Create a program to display a classification based on age.
//

package main

import (
	"fmt"
	random "math/rand/v2"
)

func getAge() int {
	return random.IntN(20)
}

func main() {
	//--Requirements:
	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+

	for i := 0; i <= 10; i++ {
		age := getAge()
		fmt.Println("The random age is:", age)
		switch {
		case age == 0:
			fmt.Println("newborn")
		case age < 4:
			fmt.Println("toddler")
		case 4 <= age && age <= 12:
			fmt.Println("child")
		case 13 <= age && age <= 17:
			fmt.Println("teenager")
		default:
			fmt.Println("adult")
		}
	}
}
