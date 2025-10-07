package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notesapp/note" // import the notes package (replace with actual module path
)

func main() {

	fmt.Println("Hi Marty")

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return // stop processing if error
	}

	userNote.Display()
}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)  // listens to command line input
	text, err := reader.ReadString('\n') // reads input until newline note the single quotes (run = single character, string = multiple characters)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n") // remove the newline character
	text = strings.TrimSuffix(text, "\r") // remove the newline character

	//var value string
	//fmt.Scanln(&value) // enter notifies GO enter means done  // only handles single word input
	// if value == "" {	// if no value entered, reprompt
	//	return "", errors.New("Value is required")
	//	  }
	return text
}
