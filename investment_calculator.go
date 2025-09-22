package main

// Every go program must have a package statement at the top
// The main package is a special package. It defines a standalone executable program, not a library.
// Every executable program must have a main package
// can have multiple files making up one package
// its all about organizing code
// Need packages and can export and import feattures across
// import is from the fmt package - not a package written from us
// go comes with a huge standard library tyat are part of the compiler
// you can use any name you want.
// main is a special package name telling go this will be the main entry point from the application we are building
// you will not always use go run to run your code
// typically run a go build in project folder to create an executable that can be run on systems without go installed
// module is simply a go project.  module with on package
// go mod init <module-name>  - creates a go.mod file in the root of your project
// go.mod file defines the module and its dependencies

// go mode init will turn a project into a module.  Must give module a namelk and where it can be held (e.g. from github)ccccccjeiuijlriivrjbfuhdrvfhnbigbtkhhcvdlnct
// go mod init github.com/username/projectname OR EXAMPLE.COM/first-app
// go build
// ./first-app
// package main tells go where to start the execution

// import another package

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	// future value = investment amount * (1 + expected return rate/100)^years

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)

}

// func must also be named main
// main function is the entry point for the program
// every executable program must have a main function in the main package
// most code must be in a function
// must have only one main function
