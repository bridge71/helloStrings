package services

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) AuthUser(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if user.PasswordHash == "" {
		return http.StatusForbidden, models.Message{RetMessage: "missing password"}
	}

	if user.Nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "missing nickname"}
	}

	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "error bind"}
	}
	auth := &models.User{}
	s.UserRepository.CheckUserName(c, auth, user.Nickname)

	err = bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return http.StatusNotAcceptable, models.Message{RetMessage: "wrong nickname or password"}
	}

	user.PasswordHash = ""
	return http.StatusOK, models.Message{
		RetMessage: "authentication successful",
		User:       *user,
	}
}

func (s *UserService) Test(c *gin.Context) (int, models.Message) {
	return http.StatusOK, models.Message{
		RetMessage: "test successful",
	}
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (s *UserService) CreateUser(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if user.Nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "nickname does not exist"}
	}
	if user.Email == "" {
		return http.StatusForbidden, models.Message{RetMessage: "email does not exist"}
	}
	if !emailRegex.MatchString(user.Email) {
		return http.StatusForbidden, models.Message{RetMessage: "illegal email"}
	}

	// password, f := c.GetPostForm("password")
	if user.PasswordHash == "" {
		return http.StatusForbidden, models.Message{RetMessage: "password does not exist"}
	}
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}

	user1 := &models.User{}
	user2 := &models.User{}
	s.UserRepository.CheckUserEmail(c, user1, user.Email)
	s.UserRepository.CheckUserName(c, user2, user.Nickname)

	fmt.Println(user2.Nickname)
	if user2.Nickname != "" {
		return http.StatusNotAcceptable, models.Message{RetMessage: "nickname has been occupied"}
	}
	if user1.Email != "" {
		return http.StatusNotAcceptable, models.Message{RetMessage: "email has been occupied"}
	}

	encryptedPassword, err := EncryptPassword(user.PasswordHash)
	for err != nil {
		encryptedPassword, err = EncryptPassword(user.PasswordHash)
	}
	user.PasswordHash = encryptedPassword

	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		err := s.UserRepository.CreaterUser(c, user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when insert user into database",
		}
	}

	user.PasswordHash = ""
	return http.StatusOK, models.Message{
		RetMessage: "nickname and email are acceptable",
		User:       *user,
	}
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
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
