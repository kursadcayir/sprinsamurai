package server

import (
    "net/http"
)

// StartServer starts the HTTP server
func StartServer() {
    http.ListenAndServe(":8000", nil)
}
