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
	// Map with key of "book" and value of another Map with string "date" with a simple array of two
	// They array at index 0 is "name of member" string and the index at 1 is "check out time" string
	// The array at index 2 if "check in time"
	// Book type in the key of the first map tells if currently checked in or out.
	books map[*Book][]*Status

	// Map with key shows member and
	members map[*Member][]*Book
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

func (l *Library) addBooksMember(book *Book, member *Member) {
	l.members[member] = append(l.members[member], book)
}

func (l *Library) removeBooksMember(book *Book, member *Member) {
	for index, bookRemove := range l.members[member] {
		if bookRemove == book {
			lastIndex := len(l.members[member]) - 1

			// 1. Move the last element into the gap of the one being removed
			l.members[member][index] = l.members[member][lastIndex]

			// 2. Erase the last element to prevent a pointer memory leak
			l.members[member][lastIndex] = nil

			// 3. Shrink the slice length by 1
			l.members[member] = l.members[member][:lastIndex]
			break
		}
	}
}

type Status struct {
	memberPointer *Member
	checkIn       string
	checkOut      string
}

func (l *Library) addBook(book *Book) {
	l.books[book] = []*Status{}
}

func (l *Library) removeBook(book *Book) {
	delete(l.books, book)
}

func (l *Library) addMember(member *Member) {
	l.members[member] = []*Book{}
}

func (l *Library) removeMember(member *Member) {
	delete(l.members, member)
}

func (l *Library) checkInBook(book *Book, member *Member) {
	if !book.checkedOut {
		fmt.Println("The book is available. Feel free to checkout the book. ")
	} else {
		// update book
		lastIndex := len(l.books[book]) - 1
		l.books[book][lastIndex].checkIn = time.Now().Format("2006-01-02 15:04:05")

		//update member
		l.removeBooksMember(book, member)
		book.checkedOut = false
	}
}

func (l *Library) checkOutBook(book *Book, member *Member) {
	if book.checkedOut {
		fmt.Println("Sorry, the book is currently checked out")
	} else {
		// update book
		newCheckOut := Status{memberPointer: member, checkIn: time.Now().Format("2006-01-02 15:04:05")}
		l.books[book] = append(l.books[book], &newCheckOut)

		//update member
		l.addBooksMember(book, member)
		book.checkedOut = true
	}
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

	centralLibrary := Library{books: make(map[*Book][]*Status),
		members: make(map[*Member][]*Book)}

	alan := Member{"Alan", 17, 2019}
	robert := Member{"Robert", 69, 1990}
	jennifer := Member{"Jennifer", 19, 2007}
	chloe := Member{"Chloe", 31, 2000}

	hobbit := Book{"The Hobbit", 310, "fantasy", false}
	fellowship := Book{"The Lord of the Rings: Fellowship of the Ring", 423, "fantasy", false}
	twoTowers := Book{"The Lord of the Rings: The Two Towers", 352, "fantasy", false}
	returnOfKing := Book{"The Lord of the Rings: The Return of the King", 416, "fantasy", false}
	projectHailMary := Book{"Project Hail Mary", 476, "scifi", false}
	onlyGoodIndians := Book{"The Only Good Indians", 320, "horror", false}
	bookLovers := Book{"Book Lovers", 384, "romance", false}

	// justDate := time.Now().Format(time.DateOnly)
	// justTime := time.Now().Format(time.TimeOnly)
	// combinedTime := time.Now().Format("2006-01-02 15:04:05")

	// fmt.Println(justDate, justTime, combinedTime)

	centralLibrary.addBook(&hobbit)
	centralLibrary.addBook(&fellowship)
	centralLibrary.addBook(&twoTowers)
	centralLibrary.addBook(&returnOfKing)
	centralLibrary.addBook(&projectHailMary)
	centralLibrary.addBook(&onlyGoodIndians)
	centralLibrary.addBook(&bookLovers)

	centralLibrary.addMember(&alan)
	centralLibrary.addMember(&robert)
	centralLibrary.addMember(&jennifer)
	centralLibrary.addMember(&chloe)

	fmt.Println(centralLibrary)
}
