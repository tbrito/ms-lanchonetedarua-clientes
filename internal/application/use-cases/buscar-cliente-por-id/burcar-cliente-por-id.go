package use_cases

import (
	"github.com/google/uuid"
	use_cases "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
)

type BuscarClientePorIdUseCase struct {
	repository domain.IClienteRepository
}

func NovoBuscarClientePorIdUseCase(repository domain.IClienteRepository) *BuscarClientePorIdUseCase {
	return &BuscarClientePorIdUseCase{repository}
}

func (c *BuscarClientePorIdUseCase) Execute(id uuid.UUID) (*use_cases.ClienteOutput, error) {

	cliente, err := c.repository.BuscarClientePorId(id)

	if err != nil {
		return nil, err
	}

	clienteOutput := use_cases.ClienteOutputFromModel(cliente)

	return &clienteOutput, nil
}
