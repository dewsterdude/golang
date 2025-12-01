package main

import "fmt"

// recursion - when a function calls itself

func main() {

	fact := factorial(5)

	fmt.Println(fact)

}

func factorial(number int) int {

	// must define an exit condition

	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1) // recursion
	}

	/*	result := 1

		for i := 1; i <= number; i++ {
			result = result * i
		}
		return result
	*/

}

// factorial of 5 =>  5 * 4 * 3 * 2 * 1
