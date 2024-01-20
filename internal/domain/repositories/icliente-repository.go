package domain

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
)

type IClienteRepository interface {
	Save(cliente *entities.Cliente) error
	Listar() ([]*entities.Cliente, error)
	BuscarClientePorId(uuid uuid.UUID) (*entities.Cliente, error)
}
