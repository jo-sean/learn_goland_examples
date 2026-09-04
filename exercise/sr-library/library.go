//--Summary:
//  Create a program to manage lending of library books.
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned

package main

import (
	"fmt"
	"time"
)

const (
	Tolkien = "J.R.R. Tolkien"
)

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
	author     string
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
			lastIndex := l.members[member]
			if len(lastIndex) != 0 {
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
}

type Status struct {
	memberPointer *Member
	checkIn       string
	checkOut      string
}

func (l *Library) addBook(book *Book) {
	l.books[book] = []*Status{}
}

// If implementing, must account for any member that has checked out the book
// and remove it from the member's list of checked out books.
// func (l *Library) removeBook(book *Book) {
// 	delete(l.books, book)
// }

func (l *Library) addMember(member *Member) {
	l.members[member] = []*Book{}
}

// If implementing, must account for any books they have out and have those
// immediately returned by looping all their pointerBooks on their list.
// func (l *Library) removeMember(member *Member) {
// 	delete(l.members, member)
// }

func findMemberByBook(library *Library, book *Book) *Member {
	// bookRecords check if empty to prevent runtime error
	bookRecords := library.books[book]
	if len(bookRecords) == 0 {
		return nil
	}

	// bookMemberString check if exists and return pointer to Member to prevent runtime error
	bookMemberString := bookRecords[len(bookRecords)-1]
	if _, exists := library.members[bookMemberString.memberPointer]; exists {
		return bookMemberString.memberPointer
	}
	return nil
}

func (l *Library) checkInBook(book *Book) bool {
	if !book.checkedOut {
		fmt.Println("Book is not checked out. Feel free to checkout the book.")
		return false
	} else {
		member := findMemberByBook(l, book)
		// update book
		lastIndex := l.books[book]
		if len(lastIndex) != 0 {
			lastIndex := len(l.books[book]) - 1
			timeStamp := time.Now().Format("2006-01-02 15:04:05.999999999")
			l.books[book][lastIndex].checkIn = timeStamp

			//update member relationshsip
			l.removeBooksMember(book, member)
			book.checkedOut = false

			fmt.Printf("Hello %s. You have returned %s at %s\n", member.name, book.title, timeStamp)
		}
		return true
	}
}

func (l *Library) checkOutBook(book *Book, member *Member) bool {
	if book.checkedOut {
		fmt.Println("Sorry, the book is currently checked out")
		return false
	} else {
		// update book
		timeStamp := time.Now().Format("2006-01-02 15:04:05.999999999")
		newCheckOut := Status{memberPointer: member, checkOut: timeStamp}
		l.books[book] = append(l.books[book], &newCheckOut)

		//update member relationship
		l.addBooksMember(book, member)
		book.checkedOut = true

		fmt.Printf("Hello %s. You have checked out %s at %s\n", member.name, book.title, timeStamp)
		return true
	}
}

func printWholeLibrary(l *Library, label string) {
	fmt.Printf("\n==================================================\n")
	fmt.Printf("   LIBRARY REPORT: %s\n", label)
	fmt.Printf("==================================================\n")

	// 1. Every Member and what they currently have rented
	fmt.Println("\n[1] REGISTERED MEMBERS & ACTIVE RENTALS")
	fmt.Println("--------------------------------------------------")
	if len(l.members) == 0 {
		fmt.Println(" No registered members.")
	} else {
		for member, borrowedBooks := range l.members {
			fmt.Printf("• %s (Age: %d, Joined: %d)\n", member.name, member.age, member.yearJoined)
			if len(borrowedBooks) == 0 {
				fmt.Println("    Currently renting: None")
			} else {
				fmt.Println("    Currently renting:")
				for _, book := range borrowedBooks {
					fmt.Printf("     - \"%s\" by %s\n", book.title, book.author)
				}
			}
		}
	}

	// 2. Every Book in the catalog and its current status
	fmt.Println("\n[2] CATALOGUE INVENTORY STATUS")
	fmt.Println("--------------------------------------------------")
	if len(l.books) == 0 {
		fmt.Println(" No books in inventory.")
	} else {
		for book := range l.books {
			statusStr := "Available"
			if book.checkedOut {
				statusStr = "CHECKED OUT"
			}
			fmt.Printf("• \"%s\" by %s (%s, %d pages) -> Status: [%s]\n",
				book.title, book.author, book.genre, book.pages, statusStr)
		}
	}

	// 3. Every single status/log event (historical record)
	fmt.Println("\n[3] TRANSACTION HISTORY & STATUS LOGS")
	fmt.Println("--------------------------------------------------")
	hasLogs := false
	for book, history := range l.books {
		if len(history) > 0 {
			hasLogs = true
			fmt.Printf("• Logs for \"%s\":\n", book.title)
			for i, log := range history {
				inTime := log.checkIn
				if inTime == "" {
					inTime = "STILL OUT"
				}
				fmt.Printf("  [%d] Member: %s\n", i+1, log.memberPointer.name)
				fmt.Printf("      Out: %s\n", log.checkOut)
				fmt.Printf("      In:  %s\n", inTime)
			}
		}
	}
	if !hasLogs {
		fmt.Println(" No transactions logged yet.")
	}
	fmt.Printf("==================================================\n\n")
}

func main() {

	centralLibrary := Library{books: make(map[*Book][]*Status),
		members: make(map[*Member][]*Book)}

	alan := Member{"Alan", 17, 2019}
	robert := Member{"Robert", 69, 1990}
	jennifer := Member{"Jennifer", 19, 2007}
	chloe := Member{"Chloe", 31, 2000}

	hobbit := Book{"The Hobbit", 310, "fantasy", Tolkien, false}
	fellowship := Book{"The Lord of the Rings: Fellowship of the Ring", 423, "fantasy", Tolkien, false}
	twoTowers := Book{"The Lord of the Rings: The Two Towers", 352, "fantasy", Tolkien, false}
	returnOfKing := Book{"The Lord of the Rings: The Return of the King", 416, "fantasy", Tolkien, false}
	projectHailMary := Book{"Project Hail Mary", 476, "scifi", "Andy Weir", false}
	onlyGoodIndians := Book{"The Only Good Indians", 320, "horror", "Stephen Graham Jones", false}
	bookLovers := Book{"Book Lovers", 384, "romance", "Emily Henry", false}

	printWholeLibrary(&centralLibrary, "Start")
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

	centralLibrary.checkInBook(&hobbit)
	centralLibrary.checkOutBook(&hobbit, &jennifer)
	centralLibrary.checkOutBook(&fellowship, &jennifer)
	centralLibrary.checkOutBook(&twoTowers, &jennifer)
	centralLibrary.checkOutBook(&returnOfKing, &jennifer)
	centralLibrary.checkOutBook(&bookLovers, &chloe)

	centralLibrary.checkOutBook(&hobbit, &alan)
	centralLibrary.checkInBook(&fellowship)
	centralLibrary.checkOutBook(&fellowship, &alan)

	printWholeLibrary(&centralLibrary, "Check")

	centralLibrary.checkInBook(&fellowship)

}
