package use_cases

import (
	use_cases "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/entities"
)

type ListarClientesOutput struct {
	Clientes []use_cases.ClienteOutput
}

func ListarClientesOutputFromModelList(clientes []*entities.Cliente) ListarClientesOutput {

	clientesOutput := use_cases.ClienteOutputFromModelList(clientes)

	return ListarClientesOutput{
		Clientes: clientesOutput,
	}
}
