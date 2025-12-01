package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil { // dont need to close if there was an error
		return nil, errors.New("Error opening prices file")
	}

	// READ IN THE LINES FROM THE FILE
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() { // read line, one at at time
		lines = append(lines, scanner.Text())
	}

	// CHECK FOR ERRORS
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error reading file conent:")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath) // Creates or overwrites a file
	if err != nil {
		return errors.New("failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON.")
	}
	file.Close()
	return nil
}
func New(inputPath, outputPath string) FileManager { // returns copy not pointer
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
