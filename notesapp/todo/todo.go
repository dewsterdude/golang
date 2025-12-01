package todo // packages must be in their own subdirectory

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"` // accessible,... the
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

// tags struct fields are some

func New(content string) (Todo, error) { // no pointer since its not big
	if content == "" {
		return Todo{}, errors.New("All fields are required")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save() error { // no pointer because its not updated

	fileName := "todo.json"
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(todo) // convert data to json - can pass struct in
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // already in byte format, and return the error returned if something goes wrong
}
