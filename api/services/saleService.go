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

func (s *SaleService) CheckStringLen(bookSale models.BookSale) (bool, string) {
	if len(bookSale.Title) > 36 {
		return true, "title is too long"
	}
	if len(bookSale.Author) > 36 {
		return true, "author is too long"
	}
	if len(bookSale.Course) > 36 {
		return true, "course is too long"
	}
	if len(bookSale.Profession) > 36 {
		return true, "profession is too long"
	}
	return false, ""
}

func checkStringLen(by models.BookBy) (bool, string) {
	if len(by.Key) > 36 {
		return true, "key is too long"
	}
	return false, ""
}

func (s *SaleService) BookGetBy(c *gin.Context) (int, models.Message) {
	by := &models.BookBy{}
	err := c.ShouldBindJSON(by)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	isLong, message := checkStringLen(*by)
	if isLong {
		return http.StatusForbidden, models.Message{RetMessage: message}
	}
	var bookSale []models.BookSale
	fmt.Println(*by)
	switch by.By {
	case "title":
		s.SaleRepository.BookGetTitle(c, &bookSale, by.Key)
	case "profession":
		s.SaleRepository.BookGetProfession(c, &bookSale, by.Key)
	case "course":
		s.SaleRepository.BookGetCourse(c, &bookSale, by.Key)
	case "author":
		s.SaleRepository.BookGetAuthor(c, &bookSale, by.Key)
	}
	// if by.By == "title" {
	// 	s.SaleRepository.BookGetName(c, &bookSale, by.Key)
	// } else if by.By == "course" {
	// 	s.SaleRepository.BookGetCourse(c, &bookSale, by.Key)
	// } else if by.By == "profession" {
	// 	s.SaleRepository.BookGetProfession(c, &bookSale, by.Key)
	// } else if by.By == "author" {
	// 	s.SaleRepository.BookGetAuthor(c, &bookSale, by.Key)
	// }
	return http.StatusOK, models.Message{BookSale: bookSale}
}

func (s *SaleService) BookSaleSubmit(c *gin.Context) (int, models.Message) {
	bookSale := &models.BookSale{}
	err := c.ShouldBindJSON(bookSale)
	if err != nil {
		return http.StatusForbidden, models.Message{RetMessage: "Bind error"}
	}
	isLong, message := s.CheckStringLen(*bookSale)
	if isLong {
		return http.StatusForbidden, models.Message{RetMessage: message}
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

func (s *SaleService) BookGet(c *gin.Context) (int, models.Message) {
	var bookSale []models.BookSale
	s.SaleRepository.BookGet(c, &bookSale)
	return http.StatusOK, models.Message{
		RetMessage: "ok",
		BookSale:   bookSale,
	}
}
