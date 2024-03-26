package desativar_cliente

import (
	"github.com/google/uuid"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
)

type DesativarClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoDesativarClienteUseCase(repository domain.IClienteRepository) *DesativarClienteUseCase {
	return &DesativarClienteUseCase{repository}
}

func (c *DesativarClienteUseCase) Execute(id uuid.UUID) error {

	cliente, err := c.repository.BuscarPorId(id)

	if err != nil {
		return err
	}

	cliente.DesativarCliente()

	err = c.repository.Atualizar(cliente)

	if err != nil {
		return err
	}

	return nil
}
