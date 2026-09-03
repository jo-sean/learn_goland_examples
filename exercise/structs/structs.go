//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//

//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

// * Using functions, calculate the area of a rectangle
//   - Print the results to the terminal
//   - The functions must use the rectangle structure as the function parameter
func (r Rectangle) area() {
	area := r.length * r.width
	fmt.Println(area)
}

// * Using functions, calculate the perimeter of a rectangle
//   - Print the results to the terminal
//   - The functions must use the rectangle structure as the function parameter
func (r Rectangle) perimeter() {
	perimeter := 2 * (r.length + r.width)
	fmt.Println(perimeter)
}

func (r *Rectangle) doubleSize() {
	r.length = 2 * r.length // 8 length
	r.width = 2 * r.width   // 10 width
}

//   - Create a rectangle structure containing its coordinates
//   - After performing the above requirements, double the size
//     of the existing rectangle and repeat the calculations
//   - Print the new results to the terminal
func main() {
	rectangle := Rectangle{4, 5}

	rectangle.area()
	rectangle.perimeter()

	rectanglePointer := &rectangle
	rectanglePointer.doubleSize()

	rectangle.area()
	rectangle.perimeter()
}
