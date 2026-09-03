//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	name        string
	securityTag bool
}

func activateTag(item *Item) {
	item.securityTag = true
}

func deactivateTag(item *Item) {
	item.securityTag = false
}

func checkout(shoppingCart []Item) {
	for i := range shoppingCart {
		deactivateTag(&shoppingCart[i])
	}
}

func checkoutByPointer(shoppingCart []*Item) {
	for _, item := range shoppingCart {
		deactivateTag(item)
	}
}

func slicePointerPrint(pointerCart []*Item) {
	fmt.Print("The cart (by pointer reference) has: ")
	for i := 0; i < len(pointerCart); i++ {
		if i == len(pointerCart)-1 {
			fmt.Print("and ")
			fmt.Print(*pointerCart[i])
			fmt.Println(".")
			break
		}
		fmt.Print(*pointerCart[i], ", ")
	}
}

func main() {

	keyboard := Item{"Keyboard", true}
	mouse := Item{"USB-C Mouse", true}
	laptop := Item{"Laptop", true}
	ssd := Item{"External SSD 1T", true}
	headphones := Item{"Headphones", true}
	chargingBrick := Item{"USB-C Charging Brick 120W", true}

	shoppingCart := []Item{keyboard, mouse, laptop, ssd}
	fmt.Println(shoppingCart)
	fmt.Println()

	fmt.Println("Deactive laptop security tag")
	deactivateTag(&shoppingCart[2]) // Deactive laptop security tag
	fmt.Println(shoppingCart)
	fmt.Println()

	fmt.Println("Active laptop security tag")
	activateTag(&shoppingCart[2]) // Active laptop security tag back again
	fmt.Println(shoppingCart)
	fmt.Println()

	fmt.Println("Checkout items in cart")
	checkout(shoppingCart)
	fmt.Println(shoppingCart)
	fmt.Println("##################################################################################################################")
	// Create instead a slice of pointers to the Items
	pointerCart := []*Item{&headphones, &chargingBrick}
	fmt.Println("Deactive laptop security tag")
	deactivateTag(pointerCart[0]) // Deactive laptop security tag
	slicePointerPrint(pointerCart)
	fmt.Println()

	fmt.Println("Active laptop security tag")
	activateTag(pointerCart[0]) // Active laptop security tag back again
	slicePointerPrint(pointerCart)
	fmt.Println()

	fmt.Println("Checkout items in cart")
	checkoutByPointer(pointerCart)
	slicePointerPrint(pointerCart)
}
