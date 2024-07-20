
package main

import (
	"fmt"
	"net/http"
)

// Note represents a single note
type Note struct {
	ID      int
	Title   string
	Content string
}

// Store to hold the notes
var notes []Note

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Notes API")
	})
	http.ListenAndServe(":8080", nil)
}
