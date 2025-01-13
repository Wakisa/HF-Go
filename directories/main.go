package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover() // Call revover and store its return value
	if p == nil {  // if recover returned nil, there is no panic...
		return // so do nothing
	}
	err, ok := p.(error) // otherwise, get hte underluing error value...
	if ok {
		fmt.Println(err) // and print it
	} else {
		panic(p) // if the panic value isn't an error, resume panicking with the same value.
	}
}

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
	defer reportPanic()

	scanDirectory("my_directory")

}
