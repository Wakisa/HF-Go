// count tallies the number of times each line
// occurs within a file.
package main

import (
	"fmt"
	"log"

	"wakisa.com/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++ // increment the vote count for the current candidate(key)
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}
