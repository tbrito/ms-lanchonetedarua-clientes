package use_cases

import (
	"github.com/google/uuid"
	"time"
)

type AtualizarClienteInput struct {
	Id             uuid.UUID `json:"id"`
	Nome           string    `json:"nome"`
	DataNascimento time.Time `json:"data_nascimento"`
	Telefone       string    `json:"telefone"`
	Endereco       string    `json:"endereco"`
	Email          string    `json:"email"`
}
