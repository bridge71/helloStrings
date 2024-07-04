package services

import (
	"encoding/json"
	"fmt"
	"io"
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

func GetUserId(c *gin.Context) uint {
	userId, err := c.Cookie("userId")
	if err != nil {
		fmt.Println("error when getUserId")
	}
	userId64, err := strconv.ParseUint(userId, 10, 8)
	if err != nil {
		fmt.Println("error when getUserId")
	}
	return uint(userId64)
}

func GetLevel(c *gin.Context) uint {
	level, err := c.Cookie("level")
	if err != nil {
		fmt.Println("error when getLevel")
	}
	level64, err := strconv.ParseUint(level, 10, 8)
	if err != nil {
		fmt.Println("error when getLevel")
	}
	return uint(level64)
}

func GetNickname(c *gin.Context) string {
	nickname, err := c.Cookie("nickname")
	if err != nil {
		fmt.Println("error when getNickname")
	}
	return nickname
}

func (s *UserService) StoreIP(c *gin.Context) (int, models.Message) {
	ip := &models.IP{}
	err := c.ShouldBindJSON(ip)
	fmt.Println("ip", ip)
	if err != nil {
		fmt.Println(err)
		return http.StatusForbidden, models.Message{RetMessage: "error bind at ip"}
	}
	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		err := s.UserRepository.StoreIP(c, ip)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when insert ip into database",
		}
	}

	return http.StatusOK, models.Message{
		RetMessage: "ip is stored",
	}
}

func (s *UserService) Login(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "error bind at user"}
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

	cookie, err := c.Cookie("isLogin")
	if err != nil || cookie != "ture" {
		cookieUserId := strconv.FormatUint(uint64(auth.UserId), 10)
		c.SetCookie("userId", cookieUserId, 24*3600, "/", "localhost", false, true)
		cookieLevel := strconv.FormatUint(uint64(auth.Level), 10)
		c.SetCookie("level", cookieLevel, 24*3600, "/", "localhost", false, true)
		c.SetCookie("nickname", auth.Nickname, 24*3600, "/", "localhost", false, true)
		c.SetCookie("isLogin", "true", 24*3600, "/", "localhost", false, true)
	}
	return http.StatusOK, models.Message{
		RetMessage: "authentication successful",
		User:       *auth,
	}
}

func (s *UserService) GetInfoUser(c *gin.Context) (int, models.Message) {
	user := &models.User{}
	user.UserId = GetUserId(c)

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
	// strTrans, _ := base64.StdEncoding.DecodeString(str)
	// filePath := "/home/bridge71/Desktop/a.jpg"
	// err := os.WriteFile(filePath, strTrans, 0666)
	// fmt.Println("sss")
	// if err != nil {
	// 	fmt.Println("sss")
	// }

	ip := &models.IP{}
	user := &models.User{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return http.StatusOK, models.Message{
			RetMessage: "test failed",
		}
	}
	err = json.Unmarshal(body, ip)
	ip.UserId = GetUserId(c)
	if err != nil {
		return http.StatusOK, models.Message{
			RetMessage: "test failed",
		}
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		return http.StatusOK, models.Message{
			RetMessage: "test failed",
		}
	}
	fmt.Println(user)
	fmt.Println(ip)
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
		err := s.UserRepository.CreateUser(c, user)
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
