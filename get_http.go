package gethttp

import (
	"log"
	"net/http"
)

// GetHTTP is the cloud fn
func GetHTTP(w http.ResponseWriter, r *http.Request) {
	author, err := getAuthor()

	if err != nil {
		log.Printf("ERROR -- %v", err)
	}

	log.Println(author)
}
