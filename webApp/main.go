package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	message := []byte("Hello, web!")
	_, err := writer.Write(message) // Add "Hello, web" to the response
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
