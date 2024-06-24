package handlers

import (
	"database/sql"
	"net/http"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := configs.DbConn()
	defer db.Close()

	var key configs.KeyOfUser
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prepare := "SELECT nickname, email FROM Users where nickname = ? or email = ?"
	stmt, err := db.Prepare(prepare)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare statment"})
		return
	}

	var keyGet configs.KeyOfUser
	err = stmt.QueryRow(key.Nickname, key.Email).Scan(&keyGet.Nickname, &keyGet.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			insert := "insert into Users(nickname, email) values(?, ?)"
			_, err2 := db.Exec(insert, key.Nickname, key.Email)
			if err2 != nil {
				c.JSON(200, gin.H{"success": "legal infomation"})
				return
			}
		} else {
			c.JSON(500, gin.H{"error": "nickname or the email is occupied"})
			return
		}
	}
}
