package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt") // Add call to getStrings

	html, err := template.ParseFiles("view.html") // Use the contents of view.html to create a new Template
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures), // Set its SignatureCount to the lenght of the signatures slice
		Signatures:     signatures,      // Sets its Signatures field to signatures slice itself.
	}
	err = html.Execute(writer, guestbook) // Write the template content to the ResponseWriter
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") // Load the contents of new.html as the text of a template
	check(err)
	err = html.Execute(writer, nil) // Write the template to the response (there's no need to insert any data in it)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")                            // Get the value of the "signature" form field
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE                     // Options for opening the file
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600)) // Open the file.
	check(err)
	_, err = fmt.Fprintln(file, signature) // Write a signature to the file, on a new line
	check(err)
	err = file.Close() // close the file
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName) // Open the file
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())
	return lines
}
