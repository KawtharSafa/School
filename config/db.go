package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {

	dataBase := "host=localhost user=postgres password=kwthr dbname=school port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dataBase), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	return db.Debug()
}
