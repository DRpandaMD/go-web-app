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

	// we want to server static files. we need to just
	// declare what it is and point to it
	staticFileDir := http.Dir("./assets/")

	//now w need to add a file handler go does some janky
	// stuff with how it handles pathing
	// so we need to rip out the prefix
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

	//now we need to use "PathPrefix" to match all routes
	// stating with /assets/ instead of the root
	// PathPrefix registers a new route with a matcher for the URL path prefix.
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return router
}

func main() {

	router := newRouter()
	http.ListenAndServe(":8080", router)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}
