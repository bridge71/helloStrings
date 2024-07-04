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

func (h *PostHandler) PostCreate(c *gin.Context) {
	code, message := h.PostService.PostCreate(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostFetch(c *gin.Context) {
	code, message := h.PostService.PostFetch(c)
	c.JSON(code, message)
}

func (h *PostHandler) PostReadTitle(c *gin.Context) {
	code, message := h.PostService.PostReadTitle(c)
	c.JSON(code, message)
}

func (h *PostHandler) ContentReadPostId(c *gin.Context) {
	code, message := h.PostService.ContentReadPostId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentReadPostId(c *gin.Context) {
	code, message := h.PostService.CommentReadPostId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentReadUserId(c *gin.Context) {
	code, message := h.PostService.CommentReadUserId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentCreate(c *gin.Context) {
	code, message := h.PostService.CommentCreate(c)
	c.JSON(code, message)
}

func (h *PostHandler) LikesChange(c *gin.Context) {
	code, message := h.PostService.LikesChange(c)
	c.JSON(code, message)
}

func (h *PostHandler) LikesReadUserId(c *gin.Context) {
	code, message := h.PostService.LikesReadUserId(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentMarkCreate(c *gin.Context) {
	code, message := h.PostService.CommentMarkCreate(c)
	c.JSON(code, message)
}

func (h *PostHandler) CommentMarkReadUserId(c *gin.Context) {
	code, message := h.PostService.CommentMarkReadUserId(c)
	c.JSON(code, message)
}
