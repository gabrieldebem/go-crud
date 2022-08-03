package repositories

import (
	"github.com/go-crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DefaultRepository struct{}

func (d DefaultRepository) db() *gorm.DB {
	dsn := "host=psql user=postgres password=root dbname=crud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
