package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// get the sizes of several web pages
	go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	time.Sleep(5 * time.Second) // Sleep for 5 seconds to allow the goroutines to finish
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
