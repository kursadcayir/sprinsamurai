package main

import (
	"net/http"
	"springsamurai/src/middleware" // Import the middleware package
	"springsamurai/src/router"     // Import the router package
	"springsamurai/src/server"
)

func main() {
	// Create the HTTP server with middleware and router
	http.Handle("/", middleware.WAFMiddleware(middleware.RuleEngineMiddleware(router.NewRouter())))

	// Start the server
	server.StartServer()
}
