package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing!\n")) // Write data to the file.
	check(err)
	err = file.Close()
	check(err)
}
