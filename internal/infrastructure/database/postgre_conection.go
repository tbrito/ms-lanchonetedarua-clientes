package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	host := "lanchonetedaruadb.c16om6u44j69.us-east-1.rds.amazonaws.com"
	port := 5432
	user := "postgres"
	password := "QE1muGg0fwsepsH"
	dbname := "cliente"
	sslMode := "require"
	timeZone := "America/Sao_Paulo"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", host, port, user, password, dbname, sslMode, timeZone)

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falha ao se conectar ao banco de dados")
	}

	return db
}
