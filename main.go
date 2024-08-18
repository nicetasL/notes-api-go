
package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Note represents a single note
type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Store to hold the notes
var (
	notes  []Note
	nextID = 1
	mutex  sync.Mutex // To ensure thread-safety
)

func main() {
	http.HandleFunc("/notes", handleNotes)
	http.ListenAndServe(":8080", nil)
}

// handleNotes processes GET and POST requests for notes
func handleNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w)
	case http.MethodPost:
		createNote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getNotes retrieves all notes
func getNotes(w http.ResponseWriter) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// createNote creates a new note
func createNote(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	note.ID = nextID
	nextID++
	notes = append(notes, note)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

