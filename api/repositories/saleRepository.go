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

func (r *SaleRepository) BookCreate(c *gin.Context, bookSale *models.BookSale) error {
	return configs.DB.WithContext(c).Create(bookSale).Error
}

func (r *SaleRepository) BookReadProfession(c *gin.Context, bookSale *[]models.BookSale, profession string) {
	configs.DB.WithContext(c).Where("profession LIKE ?", "%"+profession+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookReadAuthor(c *gin.Context, bookSale *[]models.BookSale, author string) {
	configs.DB.WithContext(c).Where("author LIKE ?", "%"+author+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookReadTitle(c *gin.Context, bookSale *[]models.BookSale, title string) {
	configs.DB.WithContext(c).Where("title LIKE ?", "%"+title+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookUpdateStatus(c *gin.Context, bookSale *models.BookSale) error {
	return configs.DB.WithContext(c).Model(bookSale).Where("created_at = ? and userID = ?", bookSale.CreatedAt, bookSale.UserId).Update("is_sold", !bookSale.IsSold).Error
}

func (r *SaleRepository) BookReadId(c *gin.Context, bookSale *[]models.BookSale, userId uint) {
	configs.DB.WithContext(c).Where("userId = ?", userId).Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookReadCourse(c *gin.Context, bookSale *[]models.BookSale, course string) {
	configs.DB.WithContext(c).Where("course LIKE ?", "%"+course+"%").Order("created_at desc").Find(bookSale)
}

func (r *SaleRepository) BookFetch(c *gin.Context, bookSale *[]models.BookSale) {
	configs.DB.WithContext(c).Order("created_at desc").Find(bookSale)
}
