package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Message(r.URL.Path[1:]))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Starting server on port %v", 8081)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Message returns a formatted message
func Message(msg string) string {
	return fmt.Sprintf("hi there, i love %s!", msg)
}
