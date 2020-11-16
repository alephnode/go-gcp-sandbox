package gethttp

import (
	"fmt"
	"log"
	"net/http"
)

// GetHTTP is the cloud fn
func GetHTTP(w http.ResponseWriter, r *http.Request) {
	name, err := getAuthorName()

	if err != nil {
		log.Printf("ERROR -- %v", err)
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}
