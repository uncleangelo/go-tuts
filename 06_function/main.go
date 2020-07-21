package main

import "fmt"

//function divide program into small and reuseable piece of code
//declare function using func
//func nameOfFunction(parameters type) return type {piece of code}
func calcRectangle(x int, y int) int {
	return x * y
}

//if parameters is same type, can specify the type only once
//eg: func calcRectangle(x, y int) int {}

func main() {
	//calling the function
	rectangle := calcRectangle(4, 5)
	fmt.Println(rectangle)

	//calling function with multiple return value
	age, yearstoRetirement := calcYearsToRetirement(1990, 2020)
	fmt.Printf("I am %d years old. \n", age)
	fmt.Printf("And I will retire in %d years.", yearstoRetirement)
}

//multiple return value
func calcYearsToRetirement(birthYear, currentYear int) (int, int) {
	age := currentYear - birthYear
	yearstoRetirement := 60 - age
	return age, yearstoRetirement
}
