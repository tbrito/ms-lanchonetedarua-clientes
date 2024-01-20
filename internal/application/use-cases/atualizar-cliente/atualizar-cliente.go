package use_cases

import (
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
	"log"
)

type AtualizarClienteUseCase struct {
	repository domain.IClienteRepository
}

func NovoAtualizarClienteUseCase(repository domain.IClienteRepository) *AtualizarClienteUseCase {
	return &AtualizarClienteUseCase{repository}
}

func (c *AtualizarClienteUseCase) Execute(atualizarClienteInput AtualizarClienteInput) error {

	cliente, err := entities.AtualizarCliente(
		atualizarClienteInput.Id,
		atualizarClienteInput.Nome,
		atualizarClienteInput.DataNascimento,
		atualizarClienteInput.Telefone,
		atualizarClienteInput.Endereco,
		atualizarClienteInput.Email,
	)

	if err != nil {
		return err
	}

	log.Println(cliente)

	err = c.repository.Atualizar(cliente)

	if err != nil {
		return err
	}

	return nil
}
