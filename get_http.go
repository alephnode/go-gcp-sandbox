package gethttp

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

// GetHTTP is the cloud fn
func GetHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	checkDbIntegration()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	log.Println("I have been called.")
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
