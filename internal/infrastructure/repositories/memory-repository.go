package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
)

type memoryClienteRepository struct {
	db []*entities.Cliente
}

func NovoInMemoryClienteRepository() *memoryClienteRepository {
	return &memoryClienteRepository{
		db: make([]*entities.Cliente, 0),
	}
}

func (r *memoryClienteRepository) Save(cliente *entities.Cliente) error {
	r.db = append(r.db, cliente)

	return nil
}

func (r *memoryClienteRepository) Listar() ([]*entities.Cliente, error) {
	return r.db, nil
}

func (r *memoryClienteRepository) BuscarClientePorId(id uuid.UUID) (*entities.Cliente, error) {

	for _, c := range r.db {
		if c.Id == id {
			return c, nil
		}
	}

	return nil, errors.New("Não encontrado")
}

//escreva uma função para excluir o cliente
