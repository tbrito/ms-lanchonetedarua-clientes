package desativar_cliente

import "github.com/google/uuid"

type DesativarClienteInput struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
