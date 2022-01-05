package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, world!")
	})


	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}
// ListenAndServe is a method exported by the http packet it allows us to start
// the web server and specify a port to listen to.  The second argument it can
// take is for HTTP/2 (read up on this), but for right now we can safely pass
// nil.
