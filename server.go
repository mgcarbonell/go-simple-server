package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}

// ListenAndServe is a method exported by the http packet it allows us to start
// the web server and specify a port to listen to.  The second argument it can
// take is for HTTP/2 (read up on this), but for right now we can safely pass
// nil.
