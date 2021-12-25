package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=postgres password=postgres port=5431 sslmode=disable database=golang-alura"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic(err.Error())
	}
}
