package main

import "fmt"

func main() {
	// array has fixed size
	// defined in '[]'
	arr := [4]int{}

	arr1 := [5]string{"Hello", "Hi", "hi", "hey", "Yo"}

	fmt.Printf("Zero value array with size 4: %v\nArray of strings: %v\n", arr, arr1)

	// slices is like array but no fixed size
	slice := []int{2, 4, 6, 8}

	fmt.Printf("slice: %v\n", slice)

	// select value of array or slice
	// array and slice are zero indexed, meaning the first element is at index zero
	// to select "hey" from arr1

	hey := arr1[3]

	// to select 4 from slice

	num4 := slice[1]

	fmt.Printf("hey from arr1: %s\nNumber 4 from slice: %d\n", hey, num4)

	// change value of array or slice

	arr1[2] = "Salam"

	fmt.Printf("hi replaced with Salam: %v\n", arr1)

}
