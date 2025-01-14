package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	// get the sizes of several web pages using channels
	pages := make(chan Page) // make a channel of Page value
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, pages) // Pass the channel to responseSize.
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages // Receive the Page
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

}

func responseSize(url string, channel chan Page) {
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
	channel <- Page{URL: url, Size: len(body)} // the size of the slice of bytes is the same as the size of the page.
}
