package main

import "fmt"

func main() {
	//'range' keyword is used to loop through array, slice or map
	//example loop through slice:
	numbers := []int{0, 1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Println("slice item", i, "is", numbers[i])
	}

	//example loop through map
	heights := map[string]int{
		"John":     179,
		"Geralt":   191,
		"Yennefer": 168,
	}

	//loop through map using keys
	for name := range heights {
		fmt.Println("height for", name, "is", heights[name])
	}

	//loop through map using values
	for _, height := range heights {
		fmt.Println(height)
	}

	//loop through key pari values
	for name, height := range heights {
		fmt.Println("Key:", name, ", Value:", height)
	}
}
