// Package http provides a set of HTTP Cloud Functions samples.
package http

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

// ParseXML is an example of parsing a text/xml request.
func ParseXML(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request", http.StatusBadRequest)
	}
	if err := xml.Unmarshal(b, &d); err != nil {
		http.Error(w, "Could not parse request", http.StatusBadRequest)
	}
	if d.Name == "" {
		d.Name = "World"
	}
	fmt.Fprintf(w, "Hello, %v!", html.EscapeString(d.Name))
}
