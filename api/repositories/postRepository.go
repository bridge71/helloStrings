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

func (r *PostRepository) PostGetTitle(c *gin.Context, post *[]models.Post, title string) {
	configs.DB.WithContext(c).Where("title LIKE ?", "%"+title+"%").Order("created_at desc").Find(post)
}

func (r *PostRepository) PostGet(c *gin.Context, post *[]models.Post) {
	configs.DB.WithContext(c).Order("created_at desc").Find(post)
}

func (r *PostRepository) PostContentGet(c *gin.Context, postContent *models.PostContent, postId uint) {
	configs.DB.WithContext(c).Where("postId = ?", postId).Find(postContent)
}
