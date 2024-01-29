package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/listar-clientes"
	"net/http"
)

func ListarClientes(ctx *gin.Context, useCase *use_cases.ListarClienteUseCase) {

	clientes, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{
				"success":  true,
				"clientes": nil,
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"sucess":   true,
			"clientes": clientes.Clientes,
		})
}
