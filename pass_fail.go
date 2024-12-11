// pass_fail report whether a grade is passing or failing
package main

import (
	"fmt"
	"log"

	"wakisa.com/keyboard"
)

func TestResult() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
