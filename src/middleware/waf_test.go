package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"springsamurai/src/middleware"
)

func TestWAFMiddleware(t *testing.T) {
	testCases := []struct {
		name           string
		requestPath    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test safe path",
			requestPath:    "/endpoint1",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Test forbidden path",
			requestPath:    "/injection",
			expectedStatus: http.StatusForbidden,
		},
	}

	// Define a handler function to be used in the test cases
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {

	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request to be passed through the middleware
			req := httptest.NewRequest("GET", tc.requestPath, nil)
			// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
			rr := httptest.NewRecorder()

			// Call the WAF middleware with the test handler function
			middleware.WAFMiddleware(middleware.RuleEngineMiddleware(http.HandlerFunc(handlerFunc))).ServeHTTP(rr, req)
			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v on endpoint %v", status, tc.expectedStatus, tc.requestPath)
			}
		})
	}
}
