package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 6, 9

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 scored higher than Quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 scored higher than Quiz 1")
	} else {
		fmt.Println("Quiz 1 and Quiz 2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable class average")
	} else {
		fmt.Println("Unacceptable class average")
	}

}
