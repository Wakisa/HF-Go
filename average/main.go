// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"wakisa.com/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//for _, number := range numbers {
	numbers = append(numbers, numbers...)
	//}
	fmt.Println(average(numbers...))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
