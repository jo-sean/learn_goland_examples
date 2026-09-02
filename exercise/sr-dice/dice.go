//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//

package main

import (
	"fmt"
	"math/rand/v2"
)

func getDiceDeets() map[string]int {
	diceMap := map[string]int{
		"diceSides":  0,
		"diceRolls":  0,
		"diceNumber": 0,
	}
	var input int

	// Validate number of dices is greater or equals to 1
	for {
		fmt.Println("How many dices are we using?")
		fmt.Scan(&input)
		if input >= 1 {
			diceMap["diceNumber"] = input
			break
		}

	}

	// Validate times being rolled is greater or equals to 1.
	for {
		fmt.Println("How many times are you going to roll your dice(s)?")
		fmt.Scan(&input)
		if input >= 1 {
			diceMap["diceRolls"] = input
			break
		}

	}

	// Validate number of sides being rolled is greater or equals to 2.
	for {
		fmt.Println("How many sides do/does the dice(s) has/have?")
		fmt.Scan(&input)
		if input >= 2 {
			diceMap["diceSides"] = input
			break
		}
	}
	return diceMap
}

// Rolls dice(s) based on the number of sides
func rollTheDice(diceMap map[string]int) int {

	if diceMap["diceNumber"] == 1 {
		return rand.IntN(diceMap["diceSides"]) + 1
	} else {
		var total int
		for dice := 1; dice <= diceMap["diceNumber"]; dice++ {

			temp := rand.IntN(diceMap["diceSides"]) + 1
			fmt.Printf("Dice# %d rolled a %d\n", dice, temp)
			total += temp
		}
		return total
	}
}

// Iterate through the rounds of dice rolls, and print the total per round
func diceRollRound(diceMap map[string]int) {
	fmt.Println("#####################################################")
	for roll := 1; roll <= diceMap["diceRolls"]; roll++ {
		fmt.Println("Roll round#", roll)
		fmt.Println("Rolling...")

		// Get the roll based on the number of dice
		total := rollTheDice(diceMap)
		checkRoll(total, diceMap["diceNumber"])
		fmt.Println("#####################################################")
	}
}

func checkRoll(total int, diceTotal int) {
	//* Print additional information in these cirsumstances:
	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd
	fmt.Printf("Your roll total is %d with %d sides each\n", total, diceTotal)
	if total == 2 && diceTotal == 2 {
		fmt.Println("###SNAKE EYES###")
	} else if total == 7 {
		fmt.Println("###LUCKY 7###")
	} else if total%2 == 0 {
		fmt.Println("***Even***")
	} else {
		fmt.Println("***Odd***")
	}
}

func main() {
	diceMap := getDiceDeets()
	diceRollRound(diceMap)
}
