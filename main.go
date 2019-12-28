package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler).Methods("GET")
	return router
}

func main() {

	router := newRouter()
	http.ListenAndServe(":8080", router)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}