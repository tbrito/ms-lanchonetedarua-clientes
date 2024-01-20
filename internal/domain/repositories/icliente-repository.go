package domain

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
)

type IClienteRepository interface {
	Criar(cliente *entities.Cliente) error
	Listar() ([]*entities.Cliente, error)
	BuscarPorId(uuid uuid.UUID) (*entities.Cliente, error)
	Atualizar(cliente *entities.Cliente) error
}
