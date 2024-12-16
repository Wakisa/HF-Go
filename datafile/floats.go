// Package datafile allowas reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of file.
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName) // open the provied file name
	if err != nil {
		return numbers, err
	}
	i := 0 // This variable will track which array index we should assign to
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
