package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getuserFloat64Value(prompt string) (float64, error) {
	var userValue float64

	fmt.Println(prompt)
	fmt.Scan(&userValue)

	if userValue <= 0 {
		return 0, errors.New("Value must be greater than zero")
	}

	return userValue, nil
}

// exercise to get user input
func getUserInput() (revenue float64, expenses float64, taxRate float64) {

	fmt.Print("\nEnter Revenue: ") // print the prompt then get user input
	fmt.Scan(&revenue)

	fmt.Print("\nEnter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("\nEnter Tax Rate: ")
	fmt.Scan(&taxRate)
	return
}

const profitCalcFile string = "profitcalc.txt"

func writeAmountToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Earnings Before Tax: %0.2f \nProfit After Tax: %0.2f \nProfit Margin: %0.2f %% \n", ebt, profit, ratio)

	os.WriteFile(profitCalcFile, []byte(results), 0644)
	fmt.Printf("Writing profit calculation to file: %s \n", profitCalcFile)
	return
}

func getAmountFromFile() (float64, error) {
	data, err := os.ReadFile(profitCalcFile) // errors dont crash and return an empty byte slice
	if err != nil {                          // abscence of a useful value
		return 0, errors.New("Error reading profit calc file")
	}
	amountText := string(data)
	amount, err2 := strconv.ParseFloat(amountText, 64)
	if err2 != nil { // abscence of a useful value
		return 0, errors.New("Error converting amount to number")
	}

	return amount, nil
}

func calculateSales(revenue float64, expenses float64, taxRate float64) (profitBeforeTax float64, tax float64, profitAfterTax float64, ratio float64) {
	profitBeforeTax = revenue - expenses
	profitAfterTax = profitBeforeTax * (1 - taxRate/100)
	tax = profitBeforeTax - profitAfterTax
	ratio = profitAfterTax / revenue * 100
	return
}

func main() {

	//	fmt.Print("Enter Revenue: ") // print the prompt then get user input
	//	fmt.Scan(&revenue)
	//	fmt.Print("Enter Expenses: ")
	//	fmt.Scan(&expenses)
	//	fmt.Print("Enter Tax Rate: ")
	//	fmt.Scan(&taxRate)

	//	revenue, expenses, taxRate = getUserInput()

	revenue, err := getuserFloat64Value("Enter Revenue: ")
	expenses, err2 := getuserFloat64Value("Enter Expenses: ")
	taxRate, err3 := getuserFloat64Value("Enter Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	if err2 != nil {
		fmt.Println(err)
		return
	}
	if err3 != nil {
		fmt.Println(err)
		return
	}
	profitBeforeTax, tax, profitAfterTax, ratio := calculateSales(revenue, expenses, taxRate)

	fmt.Println("Revenue: ", revenue)
	fmt.Println("Expenses: ", expenses)
	fmt.Println("Tax Rate: ", taxRate)

	fmt.Println("Profit Before Tax: ", profitBeforeTax)
	fmt.Printf("Tax: %0.2f \n", tax)
	fmt.Println("Profit After Tax: ", profitAfterTax)
	fmt.Println("Profit Margin: ", ratio)

	writeAmountToFile(profitBeforeTax, profitAfterTax, ratio)
}
