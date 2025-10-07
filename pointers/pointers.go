package main

import "fmt"

func main() {
	age := 32

	var agePointer *int // pointer variable that can hold the memory address of an int
	agePointer = &age   // storing mem address of age in a variable

	fmt.Println("age:", age)
	fmt.Println("age:", &age) // referencing the memory address
	fmt.Println("age:", agePointer)
	fmt.Println("age:", *agePointer) // dereferencings the pointer

	// adultYears := getAdultYears(&age) // pass pointer to age
	editAgeToAdultYears(&age) // pass pointer to age

	fmt.Println("adultYears:", age)
}

func editAgeToAdultYears(age *int) {
	//	return *age - 18
	*age = *age - 18
}
