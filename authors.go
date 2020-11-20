package gethttp

import (
	"database/sql"
	"log"
)

type author struct {
	authorName, authorID string
}

func getAuthor() (author, error) {

	dbPool, err := setupDbConnectionPool()

	if err != nil {
		log.Fatalf("%s", err)
	}

	var authorResponse author
	sqlStatement := `SELECT * FROM authors WHERE author_id=3;`

	row := dbPool.QueryRow(sqlStatement)

	switch err = row.Scan(&authorResponse.authorID, &authorResponse.authorName); err {
	case sql.ErrNoRows:
		log.Println("ERROR -- No rows were returned!")
	case nil:
		log.Println("INFO -- connected to authors table")
		log.Printf("author name is %q", authorResponse)
	default:
		log.Printf("ERROR -- %q", err)
		panic(err)
	}

	return authorResponse, err
}
