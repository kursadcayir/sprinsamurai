package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// APICallMiddleware is a middleware that makes an API call after all other middleware functions have processed the request
func APICallMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer to capture the status code
		responseWriter := NewStatusCaptureResponseWriter(w)

		// Execute all previous middleware functions first
		next.ServeHTTP(responseWriter, r)

		// Extract endpoint information from the request
		endpoint := r.URL.Path
		fmt.Println("Endpoint:", endpoint)

		// Check if the response status code is forbidden
		if responseWriter.StatusCode() == http.StatusForbidden {
			// Log the forbidden status and skip the API call
			log.Println("Skipping API call due to forbidden status")
			return
		}

		// Make API call
		apiResponse, err := makeAPIRequest(endpoint)
		if err != nil {
			// Handle API request error
			http.Error(w, "API request failed", http.StatusInternalServerError)
			return
		}
		defer apiResponse.Body.Close()

		// Write API response status code to the client
		w.WriteHeader(apiResponse.StatusCode)

		// Copy API response body to the client response
		if _, err := io.Copy(w, apiResponse.Body); err != nil {
			// Handle copy error
			http.Error(w, "Failed to write API response", http.StatusInternalServerError)
			return
		}
	})
}

// StatusCaptureResponseWriter is a custom response writer to capture the status code
type StatusCaptureResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewStatusCaptureResponseWriter creates a new StatusCaptureResponseWriter
func NewStatusCaptureResponseWriter(w http.ResponseWriter) *StatusCaptureResponseWriter {
	return &StatusCaptureResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, // Default status code
	}
}

// WriteHeader captures the status code before writing headers
func (w *StatusCaptureResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// StatusCode returns the captured status code
func (w *StatusCaptureResponseWriter) StatusCode() int {
	return w.statusCode
}

// makeAPIRequest is a helper function to make the API request
func makeAPIRequest(endpoint string) (*http.Response, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "http://localhost:8081"+endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add any headers or parameters to the request if needed
	// For example:
	// req.Header.Add("Authorization", "Bearer <token>")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
