package services

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "error bind"}
	}

	isLong, message := s.CheckStringLen(*user)
	if isLong {
		return http.StatusForbidden, models.Message{RetMessage: message}
	}

	if user.PasswordHash == "" {
		return http.StatusForbidden, models.Message{RetMessage: "missing password"}
	}
	if user.Nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "missing nickname"}
	}

	auth := &models.User{}
	s.UserRepository.CheckUserName(c, auth, user.Nickname)

	err = bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return http.StatusNotAcceptable, models.Message{RetMessage: "wrong nickname or password"}
	}

	auth.PasswordHash = ""
	fmt.Println(*auth)
	return http.StatusOK, models.Message{
		RetMessage: "authentication successful",
		User:       *auth,
	}
}

func (s *UserService) GetInfoUser(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "error bind"}
	}
	auth := &models.User{}
	s.UserRepository.GetInfoUser(c, auth, user.UserId)

	auth.PasswordHash = ""
	auth.Email = strings.Split(auth.Email, "@")[0]
	return http.StatusOK, models.Message{
		RetMessage: "get information successful",
		User:       *auth,
	}
}

func (s *UserService) Test(c *gin.Context) (int, models.Message) {
	return http.StatusOK, models.Message{
		RetMessage: "test successful",
	}
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (s *UserService) CheckStringLen(user models.User) (bool, string) {
	if len(user.Nickname) > 36 {
		return true, "nickname is too long"
	}
	if len(user.PasswordHash) > 36 {
		return true, "password is too long"
	}
	if len(user.Email) > 36 {
		return true, "email is too long"
	}
	return false, ""
}

func (s *UserService) CreateUser(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	isLong, message := s.CheckStringLen(*user)
	if isLong {
		return http.StatusForbidden, models.Message{RetMessage: message}
	}
	if user.Nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "nickname does not exist"}
	}
	if user.Email == "" {
		return http.StatusForbidden, models.Message{RetMessage: "email does not exist"}
	}
	if !emailRegex.MatchString(user.Email) {
		return http.StatusForbidden, models.Message{RetMessage: "illegal email"}
	}
	qqEmail := strings.Split(user.Email, "@")[1]
	if qqEmail != "qq.com" {
		return http.StatusForbidden, models.Message{RetMessage: "not qqemail"}
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
