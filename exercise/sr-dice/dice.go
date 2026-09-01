//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//

package main

import (
	"fmt"
	random "math/rand/v2"
)

func getDiceDeets() map[string]int {
	diceMap := map[string]int{
		"diceSides":  0,
		"diceRolls":  0,
		"diceNumber": 0,
	}
	var input int
	fmt.Println("How many dices are we using?")
	fmt.Scan(&input)
	diceMap["diceNumber"] = input
	fmt.Println("How many times are you going to roll your dice(s)?")
	fmt.Scan(&input)
	diceMap["diceRolls"] = input
	fmt.Println("How many sides do/does the dice(s) has/have?")
	fmt.Scan(&input)
	diceMap["diceSides"] = input
	return diceMap
}

func rollTheDice(diceMap map[string]int) int {
	
	if diceMap["diceNumber"] == 1 {
		return random.IntN(diceMap["diceSides"] + 1)
	} else {
		var total int
		for dice := 1; dice <= diceMap["diceNumber"]; dice++ {

			temp := random.IntN(diceMap["diceSides"] + 1)
		fmt.Printf("Dice# %d is \n")
		total += temp
		}
	}
	return total
}

func main() {
	//--Requirements:
	//* Print the sum of the dice roll
	//* Print additional information in these cirsumstances:
	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd
	//* The program must handle any number of dice, rolls, and sides
	//
	//--Notes:
	//* Use packages from the standard library to complete the project
	diceMap := getDiceDeets()
	for roll := 0; roll < diceMap["diceRolls"]; roll++ {
		fmt.Println("Roll #", roll)
		fmt.Println("Rolling...")

		total (LINK THE FUNCTION. CURRENTLY HERE. ALSO, TRY TO UNDERSTAND THE ROLL TOTAL.)

		if TOTAL == 2 && diceMap["diceNumber"] == 2 {
			fmt.Println("Snake eyes")
		} else if {
			
		}
	}

	fmt.Println(diceMap)
}
