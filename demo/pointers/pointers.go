package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// The dot dereferences in the structure.
	counter.hits += 1
}

func decrement(counter *Counter) {
	// The dot dereferences in the structure.
	counter.hits -= 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	altCounter := Counter{}
	fmt.Println(altCounter)
	increment(&altCounter)
	fmt.Println(altCounter)
	decrement(&altCounter)
	fmt.Println(altCounter)

	phrase := []string{hello, world}
	fmt.Println(phrase, counter)
	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase, counter)

}
