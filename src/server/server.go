package server

import (
	"fmt"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer() {
	fmt.Printf("Server listening on port 8000\n")
	http.ListenAndServe(":8000", nil)
}
