package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	use_cases "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/atualizar-cliente"
	"net/http"
)

func AtualizarCliente(ctx *gin.Context, useCase *use_cases.AtualizarClienteUseCase, id uuid.UUID) {

	var body use_cases.AtualizarClienteInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	body.Id = id
	err := useCase.Execute(body)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}
}
