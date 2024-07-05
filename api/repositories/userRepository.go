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

func (r *UserRepository) UserCreate(c *gin.Context, user *models.User) error {
	return configs.DB.WithContext(c).Create(user).Error
}

func (r *UserRepository) IPCreate(c *gin.Context, IP *models.IP) error {
	return configs.DB.WithContext(c).Create(IP).Error
}

func (r *UserRepository) IPRead(c *gin.Context, IP *[]models.IP, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Find(IP)
}

func (r *UserRepository) UserReadId(c *gin.Context, user *models.User, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Find(user)
}

func (r *UserRepository) UserReadNickname(c *gin.Context, user *models.User, nickname string) {
	configs.DB.WithContext(c).Where("nickname = ?", nickname).Find(user)
}

func (r *UserRepository) UserReadEmail(c *gin.Context, user *models.User, email string) {
	configs.DB.WithContext(c).Where("email = ?", email).Find(user)
}
