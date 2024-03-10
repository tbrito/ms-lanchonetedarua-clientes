package remover_cliente

import "github.com/google/uuid"

type RemoverClienteInput struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
