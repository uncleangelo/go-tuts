package main

import "fmt"

func main() {
	// a := 67
	// listNum := []int{1, 2, 3, 4, 5, 6}
	// listNum = append(listNum, a)

	// fmt.Printf("%d %v \n", len(listNum), listNum)
	// i := 0
	// for i < len(listNum) {
	// 	fmt.Println(listNum[i])
	// 	i++
	// }

	// for index := range listNum {
	// 	fmt.Printf(" %d \n", index)
	// }

	// result, err := sum("9", "89")

	// if err != nil {
	// 	panic("Something went wrong")
	// }

	// fmt.Println(result)

	result := sumAll(9, 2, 4, 16, 7, 354)
	fmt.Println(result)
}

// func sum(x int, y int) (z int) {
// 	z = x + y
// 	return
// }

// func sum(x string, y string) (int, error) {
// 	a, errA := strconv.Atoi(x)
// 	b, errB := strconv.Atoi(y)
// 	if errA != nil || errB != nil {
// 		return 0, errors.New("Can't convert to number")
// 	}

// 	return a + b, nil
// }

func sumAll(x ...int) int {
	a := 0
	for _, value := range x {
		a = a + value
	}
	return a
}
