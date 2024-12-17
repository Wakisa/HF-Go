// average calculates the average of several numbers.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
