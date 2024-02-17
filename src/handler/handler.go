package handler

import (
	"fmt"
	"net/http"
)

func Endpoint1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Endpoint 1")
}

func Endpoint2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Endpoint 2")
}

func TimeSeriesHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Response from Time Series")
}
