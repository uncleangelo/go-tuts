package main

import "fmt"

//you can think struct as a food which has specific properties

//struct can be compared to class in Object Oriented Programming
//language such as C#

//a struct can be defined outside of the func main()
type employee struct {
	firstName string
	age       int
}

func main() {
	//initialize struct with field names
	abu := employee{
		age:       45,
		firstName: "Abu",
	}
	fmt.Println(abu.firstName, "is", abu.age, "years old")

	//initialize struct wihtout declaring field names
	//need to follow the order of the field names in the struct type
	ahmad := employee{"Ahmad", 30}
	fmt.Println(ahmad.firstName, "is", ahmad.age, "years old")

	//can also initialize struct without using all the field names
	//it will output the undeclared field names as zero value
	along := employee{
		firstName: "Along",
	}
	fmt.Println(along.firstName, "is a new employee here. Age is", along.age, "(zero value)")
	//the other field can be added later
	along.age = 25
	fmt.Println(along.firstName, "is", along.age, "years old")

	//method in struct
	alongToRetire := along.yearsUntilRetirement()
	fmt.Println(alongToRetire)
}

//adding a method into a struct
func (e employee) yearsUntilRetirement() int {
	return 60 - e.age
}
