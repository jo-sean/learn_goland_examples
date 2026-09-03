//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.

package main

import (
	"fmt"
)

// - Products must include the price and the name
type Product struct {
	name  string
	price float32
}

func totalListCost(shoppingList *[4]Product) {
	var total float32
	for _, product := range shoppingList {
		total += product.price
	}

	fmt.Println(total)
}

func addProducts(shoppingList *[4]Product, productList ...Product) {
	for _, newProduct := range productList {
		for i, value := range *shoppingList {
			if value.name == "" && value.price == 0 {
				shoppingList[i] = newProduct
				break
			}
		}
	}
}

func totalNumberOfItems(shoppingList *[4]Product) {
	fmt.Println("There are a total of", len(*shoppingList), "items on your shopping list.")
}

func lastItemOnList(shoppingList *[4]Product) {
	var lastIndex int
	for i, value := range *shoppingList {
		if value.name == "" && value.price == 0 {
			break
		}
		lastIndex = i
	}
	fmt.Println("The last item on your shopping list is", shoppingList[lastIndex].name+".")
}

//   - Using an array, create a shopping list with enough room
//     for 4 products
//   - Insert 3 products into the array
//   - Print to the terminal:
//   - The last item on the list
//   - The total number of items
//   - The total cost of the items
//   - Add a fourth product to the list and print out the
//     information again
func main() {

	var shoppingList [4]Product
	fmt.Println(shoppingList)

	bacon := Product{"Egg", 2.50}
	juice := Product{"Orange Juice", 4.50}
	butter := Product{"Butter", 3.50}
	cookies := Product{"Oatmeal Chocolate Chip Cookies", 5}

	addProducts(&shoppingList, bacon, juice, butter)
	fmt.Println(shoppingList)

	lastItemOnList(&shoppingList)
	totalNumberOfItems(&shoppingList)
	totalListCost(&shoppingList)

	addProducts(&shoppingList, cookies)
	fmt.Println(shoppingList)

	lastItemOnList(&shoppingList)
	totalNumberOfItems(&shoppingList)
	totalListCost(&shoppingList)

}
