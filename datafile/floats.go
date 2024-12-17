// Package datafile allowas reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64          /// a slide data type with nil by default
	file, err := os.Open(fileName) // open the provied file name
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
