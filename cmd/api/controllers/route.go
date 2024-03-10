package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	use_cases_atualizar "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/atualizar_cliente"
	use_cases_buscar_por_id "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/buscar_cliente_por_id"
	use_cases_criar "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/criar_cliente"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/desativar_cliente"
	use_cases_listar "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/listar_clientes"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/remover_cliente"
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
		CriarCliente(ctx, useCase)
	})

	clienteRoutes.GET("/", func(ctx *gin.Context) {
		useCase := use_cases_listar.NovoListarClienteUseCase(clienteRepository)
		ListarClientes(ctx, useCase)
	})

	clienteRoutes.GET("/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de UUID inválido"})
			return
		}
		useCase := use_cases_buscar_por_id.NovoBuscarClientePorIdUseCase(clienteRepository)
		BuscarClientePorId(ctx, useCase, id)
	})

	clienteRoutes.PUT("/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de UUID inválido"})
			return
		}
		useCase := use_cases_atualizar.NovoAtualizarClienteUseCase(clienteRepository)
		AtualizarCliente(ctx, useCase, id)
	})

	clienteRoutes.PATCH("/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de UUID inválido"})
			return
		}
		useCase := desativar_cliente.NovoDesativarClienteUseCase(clienteRepository)
		DesativarCliente(ctx, useCase, id)
	})

	clienteRoutes.DELETE("/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de UUID inválido"})
			return
		}
		useCase := remover_cliente.NovoRemoverClienteUseCase(clienteRepository)
		RemoverCliente(ctx, useCase, id)
	})
}
