package gethttp

import (
	"database/sql"
	"log"
)

func getAuthorName() (string, error) {

	dbPool, err := setupDbConnectionPool()

	if err != nil {
		log.Fatalf("%s", err)
	}

	var authorName string
	authorID := 3
	sqlStatement := `SELECT author_name FROM authors WHERE author_id=$1;`

	row := dbPool.QueryRow(sqlStatement, authorID)

	switch err = row.Scan(&authorName); err {
	case sql.ErrNoRows:
		log.Println("ERROR -- No rows were returned!")
	case nil:
		log.Println("INFO -- connected to authors table")
		log.Printf("author name is %q", authorName)
	default:
		log.Printf("ERROR -- %q", err)
		panic(err)
	}

	return authorName, err
}
