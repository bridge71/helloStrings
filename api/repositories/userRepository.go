package repositories

import (
	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(c *gin.Context, user *models.User) error {
	return configs.DB.WithContext(c).Create(user).Error
}

func (r *UserRepository) StoreIP(c *gin.Context, IP *models.IP) error {
	return configs.DB.WithContext(c).Create(IP).Error
}

// func (r *UserRepository) InjectAuth(c *gin.Context, userauth *models.UserAuth) error {
// 	return configs.DB.WithContext(c).Create(&userauth).Error
// }

func (r *UserRepository) GetInfoUser(c *gin.Context, user *models.User, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Find(user)
}

func (r *UserRepository) CheckUserName(c *gin.Context, user *models.User, nickname string) {
	configs.DB.WithContext(c).Where("nickname = ?", nickname).Find(user)
}

func (r *UserRepository) CheckUserEmail(c *gin.Context, user *models.User, email string) {
	configs.DB.WithContext(c).Where("email = ?", email).Find(user)
}
