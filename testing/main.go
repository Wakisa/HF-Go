package main

import (
	"fmt"

	"wakisa.com/testing/pack"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", pack.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", pack.JoinWithCommas(phrases))
}
