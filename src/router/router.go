package router

import (
	"net/http"
	"springsamurai/src/handler"
)

// NewRouter creates a new instance of the HTTP router
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/endpoint1", handler.Endpoint1Handler)
	mux.HandleFunc("/endpoint2", handler.Endpoint2Handler)
	mux.HandleFunc("/timeseries", handler.TimeSeriesHandler)
	// Add more endpoints as needed
	return mux
}
