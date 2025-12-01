package main

import "fmt"

/*. doesnt work well
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(floatMap)
}
*/

func main() {
	// string = type, 2 is allocated, 5 is max capacity) - that takes 5 items withe 2 empty slots
	// this allocates the memory to unser enough memory is allocated
	// make is useful if you know how many items you will need to support
	// might be certain optimization for certain applciations

	userNames := make([]string, 2, 5) // slice of strings + initial length of slice extra argument
	// go will create something with the size of 2

	userNames[0] = "Julie"

	// appends to the end
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	fmt.Println(courseRatings)

	for index, value := range userNames { // go thru all the userNames
		fmt.Println("INdex: ", index)
		fmt.Println("value: ", value)

	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)

	}
}
