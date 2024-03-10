package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/remover_cliente"
	"net/http"
)

func RemoverCliente(ctx *gin.Context, useCase *remover_cliente.RemoverClienteUseCase, id uuid.UUID) {

	err := useCase.Execute(id)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusNoContent,
		gin.H{
			"sucess": true,
		})
}
