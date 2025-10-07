package main

// can have multiple files making up one package
// all files in a package must have the same package name
// package name is usually the same as the directory name
// package main is a special package that tells the Go compiler that the package should compile as an executable program instead of a shared library
// every executable program must have a main package

import (
	"fmt"
	"os"

	fileOps "examples.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile string = "balance.txt"

func main() {
	//	var accountBalance float64 = 1000.00
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println(err)
	}
	//	data, err :=

	os.ReadFile("balance.txt")

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Random Phone Number", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice) // pointer to choice variable
		fmt.Println("You chose option", choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %0.2f \n", accountBalance)
		case 2:
			{
				var depositAmount float64
				fmt.Println("Enter amount to deposit:")
				fmt.Scan(&depositAmount)
				if depositAmount > 0 {
					accountBalance += depositAmount
					fmt.Printf("Your balance is: %0.2f \n", accountBalance)
					fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
				} else {
					fmt.Println("Deposit amount must be positive")
				}
			}
		case 3:
			{
				var withdrawlAmount float64
				fmt.Println("Enter amount to withdraw:")
				fmt.Scan(&withdrawlAmount)
				if withdrawlAmount > 0 && withdrawlAmount <= accountBalance {
					fmt.Printf("You withdrew $%.2f\n", withdrawlAmount)
					accountBalance -= withdrawlAmount
					fmt.Printf("Your balance is: %0.2f \n", accountBalance)
					fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
				} else {
					fmt.Println("Withdrawl amount must be positive and less than or equal to your balance")
				}

			}
		case 4:
			{
				fmt.Println("Thank you for using Go Bank!")
				return
			}
		default:
			fmt.Println("Invalid choice")
		}

	}
}

// panic - crashes the program
// error - a value that can be returned and handled gracefully by the program

/*
	if choice == 1 {
		fmt.Printf("Your balance is: %0.2f \n", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Println("Enter amount to deposit:")
		fmt.Scan(&depositAmount)
		if depositAmount > 0 {
			accountBalance += depositAmount
			fmt.Printf("Your balance is: %0.2f \n", accountBalance)
		} else {
			fmt.Println("Deposit amount must be positive")
		}
	} else if choice == 3 {
		var withdrawlAmount float64
		fmt.Println("Enter amount to withdraw:")
		fmt.Scan(&withdrawlAmount)
		fmt.Printf("You withdrew $%.2f\n", withdrawlAmount)
		accountBalance -= withdrawlAmount
		fmt.Printf("Your balance is: %0.2f \n", accountBalance)
		continue
	} else if choice == 4 {
		break // breaks out of the loop
	} else {
		fmt.Println("Invalid choice")
	}
*/
