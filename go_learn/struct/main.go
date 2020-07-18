package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type male struct {
	person
	beard bool
}

func main() {
	p1 := person{
		"Amad",
		"Mamad",
		32,
	}

	p2 := person{
		firstName: "Ben",
		lastName:  "Ten",
		age:       10,
	}

	p2.age = 55

	// fmt.Println(p1.firstName, p2.age)
	p1Greet := p1.greet("Hi")
	p2Greet := p2.greet("Hello")
	fmt.Println(p1Greet)
	fmt.Println(p2Greet)
}

func (p person) greet(s string) string {
	return s + ", my name is " + p.firstName
}
