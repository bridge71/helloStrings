package cofnigs

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	// dbPass := "password"
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
