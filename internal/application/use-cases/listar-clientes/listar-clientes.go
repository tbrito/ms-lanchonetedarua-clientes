package use_cases

import (
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
)

type ListarClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoListarClienteUseCase(repository domain.IClienteRepository) *ListarClienteUseCase {
	return &ListarClienteUseCase{repository}
}

func (c *ListarClienteUseCase) Execute() (*ListarClientesOutput, error) {
	clientes, err := c.repository.Listar()

	if err != nil {
		return nil, err
	}

	clientesOutput := ListarClientesOutputFromModelList(clientes)

	return &clientesOutput, nil
}
