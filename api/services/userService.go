package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CheckUser(c *gin.Context) (int, models.Message) {
	nickname, f1 := c.GetPostForm("nickname")
	email, f2 := c.GetPostForm("email")
	fmt.Println(email)
	fmt.Println(nickname)

	if !f1 || nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "nickname does not exist"}
	}
	if !f2 || email == "" {
		return http.StatusForbidden, models.Message{RetMessage: "email does not exist"}
	}
	user := &models.User{
		Nickname: nickname,
		Email:    email,
		// Level:    0,
	}
	user1 := &models.User{}
	user2 := &models.User{}
	s.UserRepository.CheckUserEmail(c, user1, user.Email)
	s.UserRepository.CheckUserName(c, user2, user.Nickname)

	if user1.Email != "" {
		return http.StatusNotAcceptable, models.Message{RetMessage: "email has been occupied"}
	}
	if user2.Nickname != "" {
		return http.StatusNotAcceptable, models.Message{RetMessage: "nickname has been occupied"}
	}

	err := s.UserRepository.CreaterUser(c, user)
	if err != nil {
		return http.StatusInternalServerError, models.Message{RetMessage: "something unusual happened when insert into database"}
	}
	return http.StatusOK, models.Message{RetMessage: "nickname and email are acceptable"}
}

func RandString() string {
	length, _ := strconv.ParseUint(os.Getenv("PASS_LEN"), 10, 8)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	modLen, _ := strconv.ParseUint(os.Getenv("PASS_RAN"), 10, 8)
	length += configs.CustomRand.Uint64() % modLen
	str := make([]byte, length)
	lenOfCharset := len(charset)
	for i := range str {
		str[i] = charset[configs.CustomRand.Intn(lenOfCharset)]
	}
	return string(str)
}
