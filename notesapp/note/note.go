package note // packages must be in their own subdirectory

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`      // accessible,... the
	Content   string    `json:"content"`    // accessible and added struct tags for Json output
	CreatedAt time.Time `json:"created_at"` // accessible and added struct tags for Jason output
}

// tags struct fields are some

func New(title, content string) (Note, error) { // no pointer since its not big

	if title == "" || content == "" {
		return Note{}, errors.New("All fields are required")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	go fmt.Printf("Your Note titled: %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error { // no pointer because its not updated

	fileName := strings.ReplaceAll(note.Title, " ", "_") // new string not editing
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) // convert data to json - can pass struct in
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // already in byte format, and return the error returned if something goes wrong
}
