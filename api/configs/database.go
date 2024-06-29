package configs

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bridge71/helloStrings/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB           *gorm.DB
	CustomSource rand.Source
	CustomRand   *rand.Rand
)

func LoadConfigs() {
	CustomSource = rand.NewSource(time.Now().UnixNano())
	CustomRand = rand.New(CustomSource)
}

func init() {
	// dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	ip := "localhost"
	port := "3306"
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		ip,
		port,
		dbName,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto-migrate the models
	DB.AutoMigrate(&models.User{}, &models.BookSale{})
}
