package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)              // printl current directory
	files, err := os.ReadDir(path) // get a slice with the directory's contents.
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("directories")
	if err != nil {
		log.Fatal(err)
	}
}
