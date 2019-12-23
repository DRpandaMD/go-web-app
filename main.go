package main

import (
	// we will need the "fmt" pacakge for I/O
	"fmt"

	// we will need the "net/http" library to implment HTTP clients and servers
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even trat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)

	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprint(writer, "Hello World!")
}
