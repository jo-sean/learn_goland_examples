//--Summary:
//  Create a program to manage parts on an assembly line.
//

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func printAssemblyLine(assemblyLine []Part) {
	for i, value := range assemblyLine {
		fmt.Println("Item", i, "on the assembly is: ", value)
	}
}

// --Requirements:
// * Using a slice, create an assembly line that contains type Part
// * Perform the following:
//   - Create an assembly line having any three parts
//   - Add two new parts to the line
//   - Slice the assembly line so it contains only the two new parts
//   - Print out the contents of the assembly line at each step
func main() {
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Nuts"
	assemblyLine[1] = "Bolts"
	assemblyLine[2] = "Rivets"

	printAssemblyLine(assemblyLine)

	fmt.Println()

	assemblyLine = append(assemblyLine, "Washers", "Screws")
	printAssemblyLine(assemblyLine)

	fmt.Println()
	newItemsOnly := assemblyLine[3:]
	printAssemblyLine(newItemsOnly)

}
