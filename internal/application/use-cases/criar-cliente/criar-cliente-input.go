package use_cases

import "time"

type CriarClienteInput struct {
	Nome           string    `json:"nome" binding:"required"`
	DataNascimento time.Time `json:"data_nascimento"`
	Telefone       string    `json:"telefone" binding:"required"`
	Endereco       string    `json:"endereco"`
	Email          string    `json:"email" binding:"required"`
}
