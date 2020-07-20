package main

import "fmt"

func main() {
	//if else statement is to determine which one of the two(or more)
	//expression is to be executed

	//the way of writing if else statement is the same as any other
	//programming language

	scores := map[string]int{
		"Ahmad": 52,
		"Ben":   72,
		"John":  90,
	}

	if scores["Ahmad"] > scores["Ben"] {
		fmt.Println("Ahmad's score is better than Bens's")
	} else if scores["Ahmad"] < scores["Ben"] {
		fmt.Println("Bens's score is better than Ahmad's")
	} else {
		fmt.Println("Both of them are draw")
	}
}
