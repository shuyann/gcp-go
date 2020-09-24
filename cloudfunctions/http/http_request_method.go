// Package http provides a set of HTTP Cloud Functions samples.
package http

import (
	"fmt"
	"net/http"
)

// HelloHTTPMethod is an HTTP Cloud Functions.
// It uses the request method to differentiate the response.
func HelloHTTPMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Hello World!")
	case http.MethodPut:
		http.Error(w, "403 - Forbidden", http.StatusForbidden)
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
