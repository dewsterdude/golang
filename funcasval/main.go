package main

import "fmt"

func main() {

	numbers := []int{1, 10, 15}
	//	sum := sumup(numbers)

	// can write functions that work with any number of parms (variadic functions)
	sum2 := sumup(1, 10, 15, 40, 50)

	anotherSum := sumup(1, numbers...)
	//	fmt.Println(sum)
	fmt.Println(sum2)
	fmt.Println(anotherSum)
}

//func sumup(numbers []int) int {
//func sumup(startingValue int, numbers []int) int {
// func sumup(numbers ...int) int {

// below talks about passing a list of parameter values, or you can accept a list of them
// parameter lists are just variable  ... collects the variables and merge to a slice for you

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val // sum = sum + val
	}
	return sum
}
