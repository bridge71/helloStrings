package repositories

import (
	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/gin-gonic/gin"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) CreateInfo(c *gin.Context, post *models.Post) error {
	return configs.DB.WithContext(c).Create(post).Error
}

func (r *PostRepository) InsertContent(c *gin.Context, content *models.PostContent) error {
	return configs.DB.WithContext(c).Create(content).Error
}
