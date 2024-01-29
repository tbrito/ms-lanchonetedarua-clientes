package use_cases

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
	"log"
)

type CriarClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoCriarClienteUseCase(repository domain.IClienteRepository) *CriarClienteUseCase {
	return &CriarClienteUseCase{repository}
}

func (c *CriarClienteUseCase) Execute(criarClienteInput CriarClienteInput) (*uuid.UUID, error) {
	cliente, err := entities.NovoCliente(
		criarClienteInput.Nome,
		criarClienteInput.DataNascimento,
		criarClienteInput.Telefone,
		criarClienteInput.Endereco,
		criarClienteInput.Email,
	)

	if err != nil {
		return nil, err
	}

	log.Println(cliente)

	clienteId, err := c.repository.Criar(cliente)

	if err != nil {
		return nil, err
	}

	return clienteId, nil
}
