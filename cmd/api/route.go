package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tbrito/ms-lanchonetedarua-clientes/cmd/api/controllers"
	use_cases_buscar_por_id "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/buscar-cliente-por-id"
	use_cases_criar "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/criar-cliente"
	use_cases_listar "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/listar-clientes"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/infrastructure/database"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/infrastructure/repositories"
	"net/http"
)

func ClienteRoutes(router *gin.Engine) {

	clienteRoutes := router.Group("/clientes")

	db := database.NewPostgresDB()
	clienteRepository := repositories.NovoClienteRepository(db)

	clienteRoutes.POST("/", func(ctx *gin.Context) {
		useCase := use_cases_criar.NovoCriarClienteUseCase(clienteRepository)
		controllers.CriarCliente(ctx, useCase)
	})

	clienteRoutes.GET("/", func(ctx *gin.Context) {
		useCase := use_cases_listar.NovoListarClienteUseCase(clienteRepository)
		controllers.ListarClientes(ctx, useCase)
	})

	clienteRoutes.GET("/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de UUID inválido"})
			return
		}
		useCase := use_cases_buscar_por_id.NovoBuscarClientePorIdUseCase(clienteRepository)
		controllers.BuscarClientePorId(ctx, useCase, id)
	})
}
