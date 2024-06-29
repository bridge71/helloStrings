package handlers

import (
	"github.com/bridge71/helloStrings/api/services"
	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	SaleService *services.SaleService
}

func NewSaleHandler(saleService *services.SaleService) *SaleHandler {
	return &SaleHandler{SaleService: saleService}
}

func (h *SaleHandler) BookCreateSale(c *gin.Context) {
	code, message := h.SaleService.BookSaleSubmit(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookGetName(c *gin.Context) {
	code, message := h.SaleService.BookGetName(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookGetAuthor(c *gin.Context) {
	code, message := h.SaleService.BookGetAuthor(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookGetCourse(c *gin.Context) {
	code, message := h.SaleService.BookGetCourse(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookGetProfession(c *gin.Context) {
	code, message := h.SaleService.BookGetProfession(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookGet(c *gin.Context) {
	code, message := h.SaleService.BookGet(c)
	c.JSON(code, message)
}

//
// func (h *UserHandler) CheckUser(c *gin.Context) {
// 	code, message := h.UserService.CreateUser(c)
// 	c.JSON(code, message)
// }
