package use_cases

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
)

type BuscarClientePorIdUseCase struct {
	repository domain.IClienteRepository
}

func NovoBuscarClientePorIdUseCase(repository domain.IClienteRepository) *BuscarClientePorIdUseCase {
	return &BuscarClientePorIdUseCase{repository}
}

func (c *BuscarClientePorIdUseCase) Execute(id uuid.UUID) (*use_cases.ClienteOutput, error) {

	cliente, err := c.repository.BuscarPorId(id)

	if err != nil {
		return nil, err
	}

	clienteOutput := use_cases.ClienteOutputFromModel(cliente)

	return &clienteOutput, nil
}
