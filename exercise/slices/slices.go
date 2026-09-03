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

	var first Part = "First"
	var second Part = "Second"
	var third Part = "Third"
	// var fourth Part = "Fourth"
	// var fifth Part = "Fifth"

	assemblyLine[0] = first
	assemblyLine[1] = second
	assemblyLine[2] = third

	printAssemblyLine(&assemblyLine)

}
