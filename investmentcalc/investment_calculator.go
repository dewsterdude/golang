package main

import (
	"fmt"
	"math"
)

// functions only execute when called

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {

	var inflationRate float64 = 4.0 // const cant ever be changed and must have an initial value assigned

	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	return
}

func main() {
	//  var investmentAmount, years float64 = 1000, 10. -- could also assign multiple types and omit the float 64

	//	const inflationRate float64 = 4.0 // const cant ever be changed and must have an initial value assigned
	var investmentAmount float64 // will be provided
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5 // without the VAR you can do explicit assignment
	// expectedReturnRate := 5.5  -- shorthand assignment operator - only works inside a function

	// future value = investment amount * (1 + expected return rate/100)^years

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount) // & is a pointer to the variable when asking for input - scan asks for user input

	fmt.Print("Enter years: ")
	fmt.Scan(&years) // & is a pointer to the variable when asking for input - scan asks for user inpu

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate) // & is a pointer to the variable when asking for input - scan asks for user input

	//	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f \n", futureValue) // Sprintf formats a string and returns it
	fmt.Print(formattedFV)                                           // Print formatted string

	fmt.Printf("future value: %0.2f \n", futureValue)
	fmt.Printf("future real value: %0.2f \n", futureRealValue)

}

// func must also be named main
// main function is the entry point for the program
// every executable program must have a main function in the main package
// most code must be in a function
// must have only one main function
