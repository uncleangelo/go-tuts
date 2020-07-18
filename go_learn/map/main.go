package main

import "fmt"

func main() {
	scores := map[string]int{
		"Ahmad": 90,
		"Ben":   89,
		"John":  33,
	}
	scores["Miranda"] = 77
	delete(scores, "John")
	// scores := map[string]int{}

	for key, value := range scores {
		fmt.Printf("Name %s, score %d \n", key, value)
	}

	fmt.Println(scores)
}
