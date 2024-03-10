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

func (r *clienteRepository) Criar(cliente *entities.Cliente) (*uuid.UUID, error) {
	tx := r.db.Create(cliente)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &cliente.Id, nil
}

func (r *clienteRepository) Listar() ([]*entities.Cliente, error) {
	var clientes []*entities.Cliente

	tx := r.db.Find(&clientes)

	return clientes, tx.Error
}

func (r *clienteRepository) BuscarPorId(id uuid.UUID) (*entities.Cliente, error) {
	var cliente entities.Cliente

	tx := r.db.First(&cliente, id)

	return &cliente, tx.Error
}

func (r *clienteRepository) Atualizar(cliente *entities.Cliente) error {
	tx := r.db.Save(&cliente)

	return tx.Error
}

func (r *clienteRepository) Remover(id uuid.UUID) error {
	var cliente entities.Cliente

	tx := r.db.First(&cliente, id)
	if tx.Error != nil {
		return tx.Error
	}

	tx = r.db.Delete(&cliente)
	return tx.Error
}
