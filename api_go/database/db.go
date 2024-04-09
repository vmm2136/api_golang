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

func ConexaoDatabase() {
	strConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnection))

	if err != nil {
		log.Panic("Erro de conexão!")
	}
}
