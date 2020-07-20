package main

import "fmt"

func main() {
	//map is equal to 'object' in other programming language
	//has key-value pair
	heights := map[string]int{
		"John":     179,
		"Geralt":   191,
		"Yennefer": 168,
	}
	fmt.Printf("map of heights: %v\n", heights)

	//adding a new value in a map
	heights["Triss"] = 155
	fmt.Printf("added Triss's height into the map: %v\n", heights)

	//deleting a value in the map requires two arguments
	//1st argument is the map name, 2nd argument is the key
	delete(heights, "John")
	fmt.Printf("remove john height from the map: %v\n", heights)
}
