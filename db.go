package gethttp

import (
	"database/sql"
	"fmt"
	"os"

	// Needed this for side effect
	_ "github.com/jackc/pgx/v4/stdlib"
)

func setupDbConnectionPool() (*sql.DB, error) {
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

	return dbPool, err

}
