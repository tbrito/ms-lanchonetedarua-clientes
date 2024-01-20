package entities

import (
	"github.com/google/uuid"
	"time"
)

type Cliente struct {
	Id             uuid.UUID
	Nome           string
	DataNascimento time.Time
	Telefone       string
	Endereco       string
	Email          string
}

func NovoCliente(nome string, dataNascimento time.Time, telefone string, endereco string, email string) (*Cliente, error) {
	cliente := &Cliente{
		Id:             uuid.New(),
		Nome:           nome,
		DataNascimento: dataNascimento,
		Telefone:       telefone,
		Endereco:       endereco,
		Email:          email,
	}

	err := cliente.IsValid()

	if err != nil {
		return nil, err
	}

	return cliente, nil
}

func (c *Cliente) IsValid() error {
	return nil
}
