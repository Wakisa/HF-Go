package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// get the sizes of several web pages using channels
	sizes := make(chan int) // make a channel of int values
	go responseSize("https://example.com/", sizes)
	go responseSize("https://golang.org/", sizes)
	go responseSize("https://golang.org/doc", sizes)

	// There will be three sends on the channel, so do three receives
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
}

func responseSize(url string, channel chan int) {
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
	channel <- len(body) // the size of the slice of bytes is the same as the size of the page.
}
