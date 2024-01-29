package use_cases

import (
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
	"time"
)

type ClienteOutput struct {
	Id             uuid.UUID `json:"id"`
	Nome           string    `json:"nome"`
	DataNascimento time.Time `json:"data_nascimento"`
	Telefone       string    `json:"telefone"`
	Endereco       string    `json:"endereco"`
	Email          string    `json:"email"`
}

func ClienteOutputFromModel(c *entities.Cliente) ClienteOutput {
	return ClienteOutput{
		Id:             c.Id,
		Nome:           c.Nome,
		DataNascimento: c.DataNascimento,
		Telefone:       c.Telefone,
		Endereco:       c.Endereco,
		Email:          c.Email,
	}
}

func ClienteOutputFromModelList(c []*entities.Cliente) []ClienteOutput {
	var clienteOutputLista []ClienteOutput

	for _, cliente := range c {
		clienteOutput := ClienteOutputFromModel(cliente)
		clienteOutputLista = append(clienteOutputLista, clienteOutput)
	}

	return clienteOutputLista
}
