package repositories

import (
	"fmt"
	"net/http"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreaterUser(c *gin.Context) {
	nickname, f1 := c.GetPostForm("nickname")
	email, f2 := c.GetPostForm("email")
	if !f1 {
		c.String(http.StatusForbidden, "nickname not exists")
	}
	if !f2 {
		c.String(http.StatusForbidden, "email not exists")
	}
	user := models.User{
		Nickname: nickname,
		Email:    email,
		Level:    1,
	}
	fmt.Println(email)
	fmt.Println(nickname)
	err := configs.DB.WithContext(c).Create(user).Error
	if err != nil {
		c.String(http.StatusNotAcceptable, "nickname or email has been occupied")
	} else {
		c.String(http.StatusOK, "nickname and email are acceptable")
	}
}
