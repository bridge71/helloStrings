package repositories

import (
	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/models"
	"github.com/gin-gonic/gin"
)

type SaleRepository struct{}

func NewsaleRepository() *SaleRepository {
	return &SaleRepository{}
}

func (r *SaleRepository) BookSale(c *gin.Context, bookSale *models.BookSale) error {
	return configs.DB.WithContext(c).Create(bookSale).Error
}

func (r *SaleRepository) BookGetProfession(c *gin.Context, bookSale *[]models.BookSale, profession string) {
	configs.DB.WithContext(c).Where("profession LIKE ?", "%"+profession+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookGetAuthor(c *gin.Context, bookSale *[]models.BookSale, author string) {
	configs.DB.WithContext(c).Where("author = ?", author).Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookGetName(c *gin.Context, bookSale *[]models.BookSale, name string) {
	configs.DB.WithContext(c).Where("book_name LIKE ?", "%"+name+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookGetCourse(c *gin.Context, bookSale *[]models.BookSale, course string) {
	configs.DB.WithContext(c).Where("course LIKE ?", "%"+course+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookGet(c *gin.Context, bookSale *[]models.BookSale) {
	configs.DB.WithContext(c).Order("created_at desc").Find(bookSale)
}
