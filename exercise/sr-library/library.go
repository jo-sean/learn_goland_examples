//--Summary:
//  Create a program to manage lending of library books.

package main

import (
	"fmt"
	"time"
)

//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned

type Library struct {
	book   Book
	member Member
}

type Book struct {
	title      string
	pages      int
	genre      string
	checkedOut bool
}

type Member struct {
	name       string
	age        int
	yearJoined int
}

// * Perform the following:
//   - Add at least 4 books and at least 3 members to the library
//   - Check out a book
//   - Check in a book
//   - Print out initial library information, and after each change
//
// * There must only ever be one copy of the library in memory at any time
//
// --Notes:
// * Use the `time` package from the standard library for check in/out times
// * Liberal use of type aliases, structs, and maps will help organize this project
func main() {

	fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"))

}
