package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := "host=lanchonetedaruadb.czcuc4gason0.us-east-1.rds.amazonaws.com user=postgres password=QE1muGg0fwsepsH dbname=cliente port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falha ao se conectar ao banco de dados")
	}

	return db
}
