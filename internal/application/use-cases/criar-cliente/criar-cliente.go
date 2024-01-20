package use_cases

import (
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
	"log"
)

type CriarClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoCriarClienteUseCase(repository domain.IClienteRepository) *CriarClienteUseCase {
	return &CriarClienteUseCase{repository}
}

func (c *CriarClienteUseCase) Execute(criarClienteInput CriarClienteInput) error {
	cliente, err := entities.NovoCliente(
		criarClienteInput.Nome,
		criarClienteInput.DataNascimento,
		criarClienteInput.Telefone,
		criarClienteInput.Endereco,
		criarClienteInput.Email,
	)

	if err != nil {
		return err
	}

	log.Println(cliente)

	err = c.repository.Save(cliente)

	if err != nil {
		return err
	}

	return nil
}
