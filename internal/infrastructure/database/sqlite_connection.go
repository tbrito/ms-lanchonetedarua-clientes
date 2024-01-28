package database

import (
	"github.com/glebarez/sqlite"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	"gorm.io/gorm"
)

func NewSqliteDB() *gorm.DB {
	dsn := "clientes.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	CreateClientesDataBase(db)

	if err != nil {
		panic("Falha ao se conectar ao banco de dados")
	}

	return db
}

func CreateClientesDataBase(db *gorm.DB) {
	db.AutoMigrate(&domain.Cliente{})
}
