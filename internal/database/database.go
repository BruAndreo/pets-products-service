package database

import (
	"github.com/bruandreo/pets-products-service/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var dsn string = "host=localhost user=pets_svc password=pets_123 dbname=pets_store port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	Database.AutoMigrate(&domain.Product{})
	return nil
}
