package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, element := range slice {
		fmt.Println()
		fmt.Printf(`At index %d there is element "%s"`, i, element)
		fmt.Println()
		fmt.Printf("Each character for %s is", element)
		for _, char := range element {
			fmt.Printf(" %q", char)
		}
		fmt.Println()
	}
}
