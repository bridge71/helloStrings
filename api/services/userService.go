package services

import (
	"net/http"
	"os"
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

func (s *UserService) AccquireInfo(c *gin.Context) (*models.User, bool, bool) {
	nickname, f1 := c.GetPostForm("nickname")
	email, f2 := c.GetPostForm("email")

	user := &models.User{
		Nickname: nickname,
		Email:    email,
		// Level:    0,
	}
	return user, f1, f2
}

func (s *UserService) AuthUser(c *gin.Context) (int, models.Message) {
	user, f1, f2 := s.AccquireInfo(c)
	if !f1 && !f2 {
		return http.StatusForbidden, models.Message{RetMessage: "misssing nickname or email"}
	}
	userAccess := &models.User{}
	if f1 {
		configs.DB.Where("nickname = ?", user.Nickname).First(userAccess)
	} else {
		configs.DB.Where("email = ?", user.Email).First(userAccess)
	}

	password, f := c.GetPostForm("password")
	if !f {
		return http.StatusForbidden, models.Message{RetMessage: "missing password"}
	}

	auth := &models.UserAuth{}
	configs.DB.Where("userId = ?", userAccess.UserId).First(auth)

	err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(password))
	if err != nil {
		return http.StatusNotAcceptable, models.Message{RetMessage: "wrong password"}
	}
	return http.StatusOK, models.Message{RetMessage: "authentication successful"}
}

func (s *UserService) CreateUser(c *gin.Context) (int, models.Message) {
	user, f1, f2 := s.AccquireInfo(c)
	if !f1 || user.Nickname == "" {
		return http.StatusForbidden, models.Message{RetMessage: "nickname does not exist"}
	}
	if !f2 || user.Email == "" {
		return http.StatusForbidden, models.Message{RetMessage: "email does not exist"}
	}

	password, f := c.GetPostForm("password")
	if !f {
		return http.StatusForbidden, models.Message{RetMessage: "password does not exist"}
	}
	encryptedPassword, err := EncryptPassword(password)
	for err != nil {
		encryptedPassword, err = EncryptPassword(password)
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

	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		err := s.UserRepository.CreaterUser(c, user)
		if err != nil {
			return err
		}

		auth := &models.UserAuth{
			UserId:       user.UserId,
			PasswordHash: encryptedPassword,
		}
		err = s.UserRepository.InjectAuth(c, auth)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{RetMessage: "something unusual happened when insert user or auth into database"}
	}

	return http.StatusOK, models.Message{RetMessage: "nickname and email are acceptable"}
}

// EncryptPassword encrypts the given password using bcrypt.
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
