package gethttp

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Needed this for side effect
	_ "github.com/jackc/pgx/v4/stdlib"
)

func checkDbIntegration() {

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

	dbPool, dbError := sql.Open("pgx", dbURI)

	if dbError != nil {
		log.Fatalf("%s", dbError)
	}

	// defer dbPool.Close()

	erryore := dbPool.Ping()
	if erryore != nil {
		log.Println("ERROR -- could not establish connection")
		panic(erryore)
	}

	sqlStatement := `SELECT author_name FROM authors WHERE author_id=$1;`
	authorID := 1
	var authorName string

	row := dbPool.QueryRow(sqlStatement, authorID)

	switch errp := row.Scan(&authorName); errp {
	case sql.ErrNoRows:
		log.Println("ERROR -- No rows were returned!")
	case nil:
		log.Println("INFO -- connected to authors table")
		log.Printf("author name is %q", authorName)
	default:
		log.Println("ERROR -- I panic, tbh")
		log.Printf("ERROR -- %q", errp)
		panic(errp)
	}

}
