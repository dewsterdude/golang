package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notesapp/note" // import the notes package (replace with actual module path
	"example.com/notesapp/todo"
)

type saver interface { // interface is a contract that guarantees that a certain value like a struct has a certain method
	Save() error // simply define a method exists with its methods and return values, and inputs
	// convention for interface that has one method,... your interface name is your method name + er
}

/*
type displayer interface {
	Display()
}
*/

type outputtable interface {
	saver
	Display()
}

/*
	type outputtable interface {
		Save() error
		Display()
	}
*/
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Save failed")
		return err
	}
	fmt.Println("Save succeeded")
	return nil
}

func main() {
	fmt.Println("Hi Marty")
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return // stop processing if error
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return // stop processing if error
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	/*
	   todo.Display()
	   err = saveData(todo)

	   	if err != nil {
	   		return
	   	}

	   userNote.Display()
	   err = saveData(userNote)

	   	if err != nil {
	   		return
	   	}
	*/
}
/*
func add(a, b interface{}) {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}
}
	*/
/* 
func printSomething(value interface{}) { // stands for any function

	// one way to decode types
/* 	intVal, ok := value.(int) // will return 2 values, the typed value and a boolean
	if ok {
		fmt.Println("integer: ", value)
		return
	}
	floatVal, ok := value.(float64) // will return 2 values, the typed value and a boolean
	if ok {
		fmt.Println("float: ", value)
		return
	}
	stringVal, ok := value.(string) // will return 2 values, the typed value and a boolean
	if ok {
		fmt.Println("string: ", value)
		return
	}
*/
	/*
	   ()

	   	switch value.(type) {
	   		case int:
	   			fmt.Println("integer: ",  value)
	   		case float64
	   			fmt.Println("float64: ", value)
	   		case string
	   			fmt.Println("string: ", value)

	   		}
	*/
}

func outputData(data outputtable) error {
	err := data.Display()
	if err != nil {
		return err
	}
	saveData(data)
	return nil

}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getNoteData() (string, string) {

	Title := getUserInput("Note title:")
	Content := getUserInput("Note content:")
	return Title, Content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)  // listens to command line input
	text, err := reader.ReadString('\n') // reads input until newline note the single quotes (run = single character, string = multiple characters)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n") // remove the newline charahp cter
	text = strings.TrimSuffix(text, "\r") // remove the newline character

	//var value string
	//fmt.Scanln(&value) // enter notifies GO enter means done  // only handles single word input
	// if value == "" {	// if no value entered, reprompt
	//	return "", errors.New("Value is required")
	//	  }
	return text
}
