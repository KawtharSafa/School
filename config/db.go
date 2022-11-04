package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() *gorm.DB {

	dataBase := "host=localhost user=postgres password=kwthr dbname=school port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dataBase), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
