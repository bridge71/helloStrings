package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func dbConn() (db *sql.DB) {
// 	dbDriver := os.Getenv("DBDRIVER")
// 	dbUser := os.Getenv("DBUSER")
// 	// dbPass := "password"
// 	dbPass := os.Getenv("DBPASS")
// 	dbName := os.Getenv("DBNAME")
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func main() {
	// router := gin.Default()
	// db := dbConn()
	// defer db.Close()
	//
	// router.GET("/", func(c *gin.Context) {
	// 	// Example of using the database
	// 	rows, err := db.Query("SELECT email from Users;")
	// 	if err != nil {
	// 		c.String(500, err.Error())
	// 		return
	// 	}
	// 	var message string
	// 	for rows.Next() {
	// 		if err := rows.Scan(&message); err != nil {
	// 			c.String(500, err.Error())
	// 			return
	// 		}
	// 	}
	// 	c.String(200, message)
	// })
	//
	router := gin.Default()
	router.Run(":8080")
}
