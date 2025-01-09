package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)              // printl current directory
	files, err := os.ReadDir(path) // get a slice with the directory's contents.
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)

		} else {
			fmt.Println(filePath)
		}
	}

}

func main() {
	scanDirectory("directories")

}
