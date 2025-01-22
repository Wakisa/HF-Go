package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message)) // convert the string to a slice of bytes
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Hello, Web!") // wrtites the string to the response.
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Salut, Web!") // wrtites the string to the response.
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Namaste, Web!") // wrtites the string to the response.
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
