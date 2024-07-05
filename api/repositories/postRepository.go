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

func (r *PostRepository) PostCreate(c *gin.Context, post *models.Post) error {
	return configs.DB.WithContext(c).Create(post).Error
}

func (r *PostRepository) CommentCreate(c *gin.Context, comment *models.Comment) error {
	return configs.DB.WithContext(c).Create(comment).Error
}

func (r *PostRepository) ContentCreate(c *gin.Context, content *models.PostContent) error {
	return configs.DB.WithContext(c).Create(content).Error
}

func (r *PostRepository) PostReadTitle(c *gin.Context, post *[]models.Post, title string) {
	configs.DB.WithContext(c).Where("title LIKE ?", "%"+title+"%").Order("created_at desc").Find(post)
}

func (r *PostRepository) PostReadNickname(c *gin.Context, post *[]models.Post, nickname string) {
	configs.DB.WithContext(c).Where("nickname = ?", nickname).Order("created_at desc").Find(post)
}

func (r *PostRepository) PostFetch(c *gin.Context, post *[]models.Post) {
	configs.DB.WithContext(c).Order("created_at desc").Find(post)
}

func (r *PostRepository) PostReadId(c *gin.Context, post *[]models.Post, ids []uint) {
	configs.DB.WithContext(c).Where(ids).Order("created_at desc").Find(post)
}

func (r *PostRepository) ContentReadPostId(c *gin.Context, postContent *models.PostContent, postId uint) {
	configs.DB.WithContext(c).Where("postId = ?", postId).Find(postContent)
}

func (r *PostRepository) CommentReadPostId(c *gin.Context, comment *[]models.Comment, postId uint) {
	configs.DB.WithContext(c).Where("postId = ?", postId).Find(comment)
}

func (r *PostRepository) CommentReadUserId(c *gin.Context, comment *[]models.Comment, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Order("created_at desc").Find(comment)
}

func (r *PostRepository) LikeCreate(c *gin.Context, like *models.Likes) error {
	return configs.DB.WithContext(c).Create(like).Error
}

func (r *PostRepository) LikeUpdateUp(c *gin.Context, post *models.Post) error {
	return configs.DB.WithContext(c).Model(post).Update("likes", post.Likes+1).Error
}

func (r *PostRepository) LikeUpdateDown(c *gin.Context, post *models.Post) error {
	return configs.DB.WithContext(c).Model(post).Update("likes", post.Likes-1).Error
}

func (r *PostRepository) LikeDel(c *gin.Context, like *models.Likes) error {
	return configs.DB.WithContext(c).Where("postId = ? and userId = ?", like.PostId, like.UserId).Delete(like).Error
}

func (r *PostRepository) LikeReadUserId(c *gin.Context, likes *[]models.Likes, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Find(likes)
}

func (r *PostRepository) LikeQuery(c *gin.Context, likes *models.Likes, userId uint, postId uint) {
	configs.DB.WithContext(c).Where("userId = ? and postId = ?", userId, postId).Find(likes)
}

func (r *PostRepository) CommentMarkCreate(c *gin.Context, commentMark *models.CommentMark) error {
	return configs.DB.WithContext(c).Create(commentMark).Error
}

func (r *PostRepository) CommentMarkReadUserId(c *gin.Context, commentMark *[]models.CommentMark, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Find(commentMark)
}
