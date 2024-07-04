package handlers

import (
	"github.com/bridge71/helloStrings/api/services"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{PostService: postService}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	code, message := h.PostService.CreatePost(c)
	c.JSON(code, message)
}

func (h *PostHandler) GetPostAll(c *gin.Context) {
	code, message := h.PostService.GetAllPost(c)
	c.JSON(code, message)
}

func (h *PostHandler) GetPostTitle(c *gin.Context) {
	code, message := h.PostService.GetPostTitle(c)
	c.JSON(code, message)
}

func (h *PostHandler) GetPostContent(c *gin.Context) {
	code, message := h.PostService.GetPostContent(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentGetPostId(c *gin.Context) {
	code, message := h.PostService.CommentGetPostId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentGetUserId(c *gin.Context) {
	code, message := h.PostService.CommentGetUserId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CreateComment(c *gin.Context) {
	code, message := h.PostService.CreateComment(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostLikesChange(c *gin.Context) {
	code, message := h.PostService.PostLikesChange(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostLikesRead(c *gin.Context) {
	code, message := h.PostService.PostLikesRead(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostCommentsAdd(c *gin.Context) {
	code, message := h.PostService.PostCommentsAdd(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostCommentsRead(c *gin.Context) {
	code, message := h.PostService.PostCommentsRead(c)
	c.JSON(code, message)
}
