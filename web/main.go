package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// get the sizes of several web pages
	responseSize("https://example.com/")
	responseSize("https://golang.org/")
	responseSize("https://golang.org/doc")
}

func responseSize(url string) {
	fmt.Println("Getting", url)    // print which URL we are retriving
	response, err := http.Get(url) // get the given URL
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body)) // the size of the slice of bytes is the same as the size of the page.
}
