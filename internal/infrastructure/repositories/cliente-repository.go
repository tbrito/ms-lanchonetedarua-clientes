package repositories

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	"gorm.io/gorm"
)

type clienteRepository struct {
	db *gorm.DB
}

func NovoClienteRepository(db *gorm.DB) *clienteRepository {
	return &clienteRepository{db}
}

func (r *clienteRepository) Save(cliente *entities.Cliente) error {
	tx := r.db.Create(cliente)

	return tx.Error
}

func (r *clienteRepository) Listar() ([]*entities.Cliente, error) {
	var clientes []*entities.Cliente

	tx := r.db.Find(&clientes)

	return clientes, tx.Error
}

func (r *clienteRepository) BuscarClientePorId(id uuid.UUID) (*entities.Cliente, error) {
	var cliente entities.Cliente

	tx := r.db.First(&cliente, id)

	return &cliente, tx.Error
}
