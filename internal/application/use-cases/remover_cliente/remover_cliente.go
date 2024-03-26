package remover_cliente

import (
	"github.com/google/uuid"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
)

type RemoverClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoRemoverClienteUseCase(repository domain.IClienteRepository) *RemoverClienteUseCase {
	return &RemoverClienteUseCase{repository}
}

func (c *RemoverClienteUseCase) Execute(id uuid.UUID) error {

	err := c.repository.Remover(id)

	if err != nil {
		return err
	}

	return nil
}
