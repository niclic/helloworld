package handlers

import (
	"net/http"
)

// Hello responds with a standard "Hello, World!" greeting.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}
