package db

import (
	"log"

	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed Connect db: ", err)
	}

	db.AutoMigrate(&models.Order{})

	return Handler{DB: db}
}
