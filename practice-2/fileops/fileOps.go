package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
	fmt.Printf("Writing value to file: %0.2f \n", value)
	return
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // errors dont crash and return an empty byte slice
	if err != nil {                    // abscence of a useful value
		return 0, errors.New("Error reading value file")
	}
	valueText := string(data)
	value, err2 := strconv.ParseFloat(valueText, 64)
	if err2 != nil { // abscence of a useful value
		return 0, errors.New("Error converting value to number")
	}

	return value, nil
}
