package main

import "fmt"

func checkNeededGroceries(shoppingList map[string]int) {

	for groceryItem, amount := range shoppingList {
		if amount > 1 {
			if amount%12 == 0 {
				fmt.Println("Need", amount/12, "dozen", groceryItem+"s")
			} else {
				fmt.Println("Need", amount, groceryItem+"s")
			}
		} else {
			fmt.Println("Need", amount, groceryItem)
		}
	}
}

func totalItemsNeeded(shoppingList map[string]int) {
	totalItemsNeeded := 0
	for _, amount := range shoppingList {
		totalItemsNeeded += amount
	}
	fmt.Println(totalItemsNeeded, "total items needed.")
}

func itemCheck(shoppingList map[string]int, groceryItem string, groceryAmount int) {
	item, found := shoppingList[groceryItem]
	fmt.Println("Need", groceryItem, "?")
	if !found {
		fmt.Println("Yup! Adding", groceryAmount, "units to the shopping list")
		shoppingList[groceryItem] = groceryAmount

	} else {
		fmt.Println("Nope, you aredy have on the list", item, "units to purchase")
	}
}

func main() {
	shoppingList := make(map[string]int)

	shoppingList["egg"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["egg"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, updated list:", shoppingList)

	checkNeededGroceries(shoppingList)

	itemCheck(shoppingList, "cereal", 4)
	checkNeededGroceries(shoppingList)

	totalItemsNeeded(shoppingList)

}
