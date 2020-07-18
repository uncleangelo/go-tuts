package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("newFile")

	if err != nil {
		log.Println("Error creating file")
	}

	file.Write([]byte("This is a new file"))
	var a int = 2

	// b := a
	// b = 77
	// fmt.Println(&a, &b)
	// fmt.Println(a, b)

	b := &a
	*b = 5
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	// dabel(&a)
	// fmt.Println(a, dabel(&a))
	name := "John"
	changeString(&name)
	fmt.Println(name)
}

// func dabel(x *int) int {
// 	*x = *x * *x
// 	return *x
// }

func changeString(s *string) {
	*s = "Changed"
}
