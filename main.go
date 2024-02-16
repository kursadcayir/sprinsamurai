package main

import (
	"net/http"
	"springsamurai/server"
)

func main() {
	// Create the HTTP server
	http.Handle("/", server.NewRouter())

	// Start the server
	server.StartServer()
}
