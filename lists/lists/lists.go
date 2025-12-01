package main

import (
	"fmt"
)

type Product struct { // data grouping mechanism
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99, 7.99}
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	fmt.Println("Prices: ", prices)

	discountPrices := []float64{101.99, 88.99}
	// prices = append(prices, discountPrices) // prices is a list of floats and discount is a list of list of floats

	prices = append(prices, discountPrices...) // ... tells go to take out the elements of the list, and use them of sepearate elements in a place like here for append
	fmt.Println("Prices + Discount: ", prices)

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.

	hobbyList := [3]string{"Guitar", "Turning", "Running"}
	fmt.Println(hobbyList)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("1. first element: ", hobbyList[0])

	newList := hobbyList[1:]
	fmt.Println("2. newList (2 & 3): ", newList)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	newSlice1 := hobbyList[:2]
	newSlice2 := hobbyList[0:2]
	fmt.Println("3. newSlice1 (1&2): ", newSlice1)
	fmt.Println("3. newSlice2: ", newSlice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	newSlice1 = hobbyList[1:]
	fmt.Println("4. newSlice1 (2&3): ", newSlice1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"UpSkill", "LearnGo"}
	fmt.Println("5. Coarse Goals (dyn): ", courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array'
	courseGoals[1] = "Coding"
	fmt.Println("6a. Coarse Goals (dyn): ", courseGoals)
	courseGoals = append(courseGoals, "Respect")
	fmt.Println("6b. Coarse Goals (dyn): ", courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	/*	productList := []Product{. // you dont have to repeat type name or titles
			Product{
				title: "phone",
				id:    "id1",
				price: 19.99,
			},
			Product{
				title: "ipad",
				id:    "id2",
				price: 199.99,
			},
		}
	*/
	productList := []Product{ // shorter notation
		{ // you dont have to repeat type name or titles
			"phone",
			"id1",
			19.99,
		},
		{
			"ipad",
			"id2",
			199.99,
		},
	}

	fmt.Println("7. Product List: ", productList)

	newproductList := append(productList,
		Product{
			title: "computer",
			id:    "id3",
			price: 27.99,
		},
	)
	fmt.Println("7. Product List: ", newproductList)

}

// Time to practice what you learned!

/*
func main() {
	// go will create another array if your array grows beyond the bound
	prices := []float64{10.99, 8.99} // omitting the number - will create an array and slice
	fmt.Println(prices[0:1])

	updatedPrices := append(prices, 5.99) // will create a NEW array w/new pointers
	prices = append(prices, 5.99)         // will redefined the prices array

	// to remove an element, you assign using slice ranging that removest he element you want

	fmt.Println("Updated: ", updatedPrices)
	fmt.Println("original slice:", prices)

	//	prices[20] = // this would fail - index out of range

}


func TemperatureData struct { // this is not ideal
	day1 float64
	day2 float64
}
*/
// array is go = array in java = lists in python
// slices are used a lot in go - will help you get a part of an array - a slice

/*
func main() {
	var productNames [4]string = [4]string{"A Book"}
	// to target individual values
	productNames[2] = "A Carpet"

	prices := [4]float64{10.99, 9.99, 45.99, 20.0} // brackets define the array and have to define how many
	// can create arrays of structs, or of other arrays

	// want a new array that omits the first and last values
	// can omit the start and write [:3] to eleminate the starting position - start at beginning then go up to some element with some index
	// can start at an element and then goto the end of the array [1:] - to get all the way to the end
	// you CANNOT use negative indexes - works ins some other programming lanugages
	// cant pick a higher end than what the array has
	// slices are like a window into an array
	// when you create an array you get access to memory
	// if you modify a slice you modifiy the element in the original array

	// arrays reuse and do not copy lengths

	//	featuredPrices := prices[1:3] // start at element at index 1, take all elements up to be excluding positing 3
/*	featuredPrices := prices[1:] // start at element at index 1, take all elements up to be excluding positing 3

	highlightedPrices := featuredPrices[:1]

	fmt.Println("Prices: ", prices)
	fmt.Println("Product Names: ", productNames)
	fmt.Println("Prices pos 3: ", prices[2])
	fmt.Println("Featured Prices: ", featuredPrices)
	fmt.Println("Highlighted Prices: ", highlightedPrices)
	featuredPrices[0] = 199.99
	fmt.Println("Featured Prices: ", featuredPrices)
	fmt.Println("Prices: ", featuredPrices)

	fmt.Println("Length of Highlighted prices: ", len(highlightedPrices), "Cap: ", cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3] // you can always select more towards end (you can reslice to right)

	// length gives you whats filled in the array, capacity tells you how much room you have
	fmt.Println("Length of featured prices: ", len(featuredPrices), "Cap: ", cap(featuredPrices))
	fmt.Println("Length of Highlighted prices: ", len(highlightedPrices), "Cap: ", cap(highlightedPrices))

}
*/
