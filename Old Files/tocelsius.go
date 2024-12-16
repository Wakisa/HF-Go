// tocelsius converts a temperature
package main

import (
	"fmt"

	"log"

	"github.com/wakisa/keyboard"
)

func ConvertTemperature() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celcius\n", celsius)
}
