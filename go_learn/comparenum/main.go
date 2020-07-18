package main

import (
	"fmt"
)

var number int = 32

// func main() {

// 	listChar := []string{"a", "b", "c"}
// 	listNum := []int{45, 77, 31}

// 	charFirst := listChar[0]
// 	numThird := listNum[2]
// 	//first := os.Args

// 	fmt.Println(charFirst)
// 	fmt.Printf("%v type %T \n", charFirst, charFirst)
// 	fmt.Println(numThird)
// 	fmt.Printf("%v type %T", numThird, numThird)
// }

// func main() {

// 	listChar := []string{"a", "b", "c", "d", "e"}

// 	huruf3 := listChar[2:]

// 	fmt.Println(huruf3)
// }

// func main() {
// 	firstArg := os.Args[1]
// 	secArg := os.Args[2]

// 	firstArgInt, _ := strconv.Atoi(firstArg)
// 	secArgInt, _ := strconv.Atoi(secArg)

// 	if firstArgInt > secArgInt {
// 		fmt.Println("First is bigger")
// 	} else if firstArgInt < secArgInt {
// 		fmt.Println("Second is bigger")
// 	} else {
// 		fmt.Println("Both are equal")
// 	}
// }

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
