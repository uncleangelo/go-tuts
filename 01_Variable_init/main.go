package main

import "fmt"

// initialize outside func, inside package scope
// only the 'var name type' can be used
var smallPositiveNumber uint8

func main() {
	// initialize variable
	var number int

	// initialize with value
	var name string = "John"

	// variable without value will default to zero
	fmt.Printf("smallPositiveNumber: %d - number: %d - name: %s\n", smallPositiveNumber, number, name)

	// shorter way to init a variable
	// must have value
	// must within func scope, can't init in package scope
	// type automatically defined by compiler
	price := 0.12

	fmt.Printf("price: %v, type(%T)\n", price, price)
}
