package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/criar-cliente"
	"net/http"
)

func CriarCliente(ctx *gin.Context, useCase *use_cases.CriarClienteUseCase) {

	var body use_cases.CriarClienteInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	clienteId, err := useCase.Execute(body)

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
		http.StatusCreated,
		gin.H{
			"sucess": true,
			"id":     clienteId,
		})
}
