package middleware

import (
	"log"
	"net/http"
	"strings"
	"sync"
)

// Rule represents a single security rule
type Rule struct {
	Name      string
	Pattern   string
	Violation string
}

// Rules is a collection of security rules
var Rules = []Rule{
	// Define your security rules here
	{
		Name:      "Test forbidden path",
		Pattern:   "/injection",
		Violation: "unsafe path detected",
	},
	{
		Name:      "Test forbidden path",
		Pattern:   "/timeseriesservice",
		Violation: "unsafe path detected",
	},
}

// RuleEngineMiddleware is a middleware that evaluates incoming requests against predefined rules
func RuleEngineMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		wg.Add(len(Rules))

		// Channel to communicate rule violations
		violationCh := make(chan struct{}, len(Rules))

		// Iterate through each rule and evaluate asynchronously
		for _, rule := range Rules {
			go func(rule Rule) {
				defer wg.Done()

				// Check if the request matches the rule pattern
				if requestMatchesPattern(r, rule.Pattern) {
					// Rule violation detected
					violationCh <- struct{}{}
				}
			}(rule)
		}

		// Wait for all rule evaluations to finish
		go func() {
			wg.Wait()
			close(violationCh)
		}()

		// Check if any rule violation detected
		select {
		case <-violationCh:
			// Rule violation detected
			log.Printf("Rule violation detected: " + r.URL.Path)
			http.Error(w, "Forbidden", http.StatusForbidden) // Example action: Return forbidden status
			return
		default:
			// No rule violation detected, call the next handler in the chain
			log.Printf("RuleEngineMiddleware passed: %s", r.URL.Path)
			next.ServeHTTP(w, r)
		}
	})
}

// requestMatchesPattern checks if the request matches the given pattern
func requestMatchesPattern(r *http.Request, pattern string) bool {
	// Implement pattern matching logic here
	// Example: Use regex.MatchString or other pattern matching logic
	log.Println("Checking request against pattern:", r.URL.Path, pattern)
	// return r.URL.Path != pattern
	return strings.Contains(r.URL.Path, pattern)
}
