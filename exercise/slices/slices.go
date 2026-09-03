//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(assemblyLine *[]Part) {
	for i, value := range *assemblyLine {
		fmt.Println(i, "- Item on the assembly is: ", value)
	}
}

func main() {
	var assemblyLine []Part

	assemblyLine[0] = "First"
	assemblyLine[1] = "Second"
	assemblyLine[2] = "Third"

	printAssemblyLine(&assemblyLine)

}
