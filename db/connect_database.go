package db

import (
	"os"
	"rbac-go/internal/auth/model"
	productmodel "rbac-go/internal/product/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	databaseConfig := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &productmodel.Product{})
	return db, nil
}
