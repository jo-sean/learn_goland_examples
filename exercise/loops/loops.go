//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//

package main

import "fmt"

func main() {
	//--Requirements:
	//* Print integers 1 to 50, except:
	//  - Print "Fizz" if the integer is divisible by 3
	//  - Print "Buzz" if the integer is divisible by 5
	//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5

	for count := 1; count <= 50; count++ {
		fmt.Println(count)
		if count%3 == 0 && count%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if count%3 == 0 {
			fmt.Println("Fizz")
		} else if count%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
