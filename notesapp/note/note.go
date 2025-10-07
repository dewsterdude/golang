package note // packages must be in their own subdirectory

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string // not accessible
	content   string // not accessible
	createdAt time.Time
}

func New(title, content string) (Note, error) { // no pointer since its not big

	if title == "" || content == "" {
		return Note{}, errors.New("All fields are required")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your Note titled: %v has the following content: \n\n%v", note.title, note.content)
}
