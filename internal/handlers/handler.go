package handlers

import "net/http"

// HTTPHandler is a function that acts as a request handler.
type HTTPHandler func(w http.ResponseWriter, r *http.Request)
