package main

// anonymous functions lecture

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	// note the func parms specified must match transformNumbers spec
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2 // anonymous function because it doesnt have name - cant be called anywhere else
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	//	transformed := transformNumbers(&numbers, ???)  // <-- error

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor // this is an anonymous factor
	} // anonymous function

}
