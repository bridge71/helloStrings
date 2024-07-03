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
