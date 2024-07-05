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

func (h *SaleHandler) BookCreate(c *gin.Context) {
	code, message := h.SaleService.BookCreate(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookReadBy(c *gin.Context) {
	code, message := h.SaleService.BookReadBy(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookFetch(c *gin.Context) {
	code, message := h.SaleService.BookFetch(c)
	c.JSON(code, message)
}

func (h *SaleHandler) BookUpdateStatus(c *gin.Context) {
	code, message := h.SaleService.BookUpdateStatus(c)
	c.JSON(code, message)
}
