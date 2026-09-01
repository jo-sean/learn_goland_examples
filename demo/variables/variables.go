package main

import "fmt"

func main() {
	var myName = "Sean"
	fmt.Println("My name is", myName)

	var username string

	fmt.Print("Type you username: ")
	fmt.Scan(&username)

	fmt.Println("Your username is now: ", username)

}
