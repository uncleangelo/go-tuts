package main

import "fmt"

//you can think struct as a food which has specific properties

//struct can be compared to class in Object Oriented Programming
//language such as C#

//a struct can be defined outside of the func main()

type Tea struct {
	name     string
	calories float64
}

func main() {
	//initialize struct with field names
	greenTea := Tea{
		calories: 2.45,
		name:     "Green Tea",
	}
	fmt.Println(greenTea)

	//initialize struct wihtout declaring field names
	//need to follow the order of the field names in the struct type
	blackTea := Tea{"Black Tea", 1}
	fmt.Println(blackTea)

	//can also initialize struct without using all the field names
	//it will output the undeclared field names as zero value
	oolong := Tea{
		name: "Oolong Tea",
	}
	fmt.Println(oolong)
}
