package services

import (
	"fmt"
	"net/http"

	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SaleService struct {
	SaleRepository *repositories.SaleRepository
}

func NewSaleService(userRepository *repositories.SaleRepository) *SaleService {
	return &SaleService{SaleRepository: userRepository}
}

func (s *SaleService) BookSaleSubmit(c *gin.Context) (int, models.Message) {
	bookSale := &models.BookSale{}
	err := c.ShouldBindJSON(bookSale)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}
	err = configs.DB.Transaction(func(tx *gorm.DB) error {
		err := s.SaleRepository.BookSale(c, bookSale)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return http.StatusInternalServerError, models.Message{
			RetMessage: "something unusual happened when insert booksale into database",
		}
	}
	return http.StatusOK, models.Message{
		RetMessage: "ok",
	}
}

func (s *SaleService) BookGetName(c *gin.Context) (int, models.Message) {
	bookRequest := models.BookSale{}
	err := c.ShouldBindJSON(&bookRequest)
	var bookSale []models.BookSale
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}
	s.SaleRepository.BookGetName(c, &bookSale, bookRequest.BookName)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}

func (s *SaleService) BookGetAuthor(c *gin.Context) (int, models.Message) {
	bookRequest := models.BookSale{}
	err := c.ShouldBindJSON(&bookRequest)
	fmt.Println("step")
	var bookSale []models.BookSale
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}
	s.SaleRepository.BookGetAuthor(c, &bookSale, bookRequest.Author)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}

func (s *SaleService) BookGetCourse(c *gin.Context) (int, models.Message) {
	bookRequest := models.BookSale{}
	err := c.ShouldBindJSON(&bookRequest)
	var bookSale []models.BookSale
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}
	s.SaleRepository.BookGetCourse(c, &bookSale, bookRequest.Course)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}

func (s *SaleService) BookGetProfession(c *gin.Context) (int, models.Message) {
	bookRequest := models.BookSale{}
	err := c.ShouldBindJSON(&bookRequest)
	var bookSale []models.BookSale
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "something error"}
	}
	s.SaleRepository.BookGetProfession(c, &bookSale, bookRequest.Profession)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}

func (s *SaleService) BookGet(c *gin.Context) (int, models.Message) {
	var bookSale []models.BookSale
	s.SaleRepository.BookGet(c, &bookSale)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}
