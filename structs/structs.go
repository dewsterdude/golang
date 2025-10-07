package main

import (
	"fmt"

	user "examples.com/structs/user"
)

func main() {
	// appuser = user{"Marty", "Dewey", "01/01/1990", time.Now()}  time.Now() returns the current time as a time struct
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// 	fmt.Println(userfirstName, userlastName, userbirthdate)

	var appUser *User
	appUser, err := user.New(userfirstName, userlastName, userbirthdate) // &user{} would be the literal way to do this, but we are using a constructor function now

	//appUser, err := New(userfirstName, userlastName, userbirthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails() // because this is a method of the struct, no parms are passed.
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

/*
	func outputUserDetails(u *user) {
		//	(*u).firstName  this is the format that would be required, but GO enables the below as a shortcut
		fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
	}
*/
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) // enter notifies GO enter means done

	return value
}
