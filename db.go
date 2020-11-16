package gethttp

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Needed this for side effect
	_ "github.com/jackc/pgx/v4/stdlib"
)

func getAuthorName() (string, error) {

	var (
		dbUser                 = os.Getenv("DB_USER")
		dbPwd                  = os.Getenv("DB_PASS")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string

	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	dbPool, err := sql.Open("pgx", dbURI)

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
