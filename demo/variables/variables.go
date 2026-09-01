package main

import "fmt"

func main() {
	var myName = "Sean"
	fmt.Println("My name is", myName)

	var username string

	fmt.Print("Type you username: ")
	fmt.Scan(&username)

	fmt.Println("Your username is now: ", username)
	var sum int
	fmt.Println("The sum is", sum) // Initialize sum as 0.

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part1 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Printf("Lesson name is %s and lesson type is %s. \n", lessonName, lessonType) // Initialize sum as 0.
}
