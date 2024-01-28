package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tbrito/ms-lanchonetedarua-clientes/cmd/api/controllers"
	use_cases4 "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases"
	use_cases5 "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/atualizar-cliente"
	use_cases2 "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/buscar-cliente-por-id"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/criar-cliente"
	use_cases3 "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/listar-clientes"
	domain "github.com/tbrito/ms-lanchonetedarua-clientes/internal/domain/repositories"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/infrastructure/database"
	"github.com/tbrito/ms-lanchonetedarua-clientes/internal/infrastructure/repositories"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type FeatureTest struct {
	Response   *httptest.ResponseRecorder
	T          *testing.T
	Router     *gin.Engine
	Repository domain.IClienteRepository
}

type Response struct {
	ID string `json:"id"`
}

var body io.Reader
var id uuid.UUID

func (ft *FeatureTest) queEuTenhoUmNovoPayloadDeCliente() error {
	body = strings.NewReader(`{"nome":"Test","dataNascimento":"2000-01-01","telefone":"123456789","endereco":"Test","email":"test@test.com"}`)
	return nil
}

func (ft *FeatureTest) euEnviarUmaRequisicaoPostPara(endpoint string) error {
	useCase := use_cases.NovoCriarClienteUseCase(ft.Repository)

	ft.Router.POST(endpoint, func(ctx *gin.Context) {
		controllers.CriarCliente(ctx, useCase)
	})

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	ft.Response = httptest.NewRecorder()

	ft.Router.ServeHTTP(ft.Response, req)

	return nil
}

func (ft *FeatureTest) oCodigoDeRespostaDeveSer(code int) error {
	assert.Equal(ft.T, code, ft.Response.Code)

	return nil
}

func (ft *FeatureTest) aRespostaDeveConterUm(field string) error {
	var response Response
	err := json.Unmarshal(ft.Response.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}
	id = uuid.MustParse(response.ID)
	assert.Contains(ft.T, ft.Response.Body.String(), field)

	return nil
}

func (ft *FeatureTest) euEnviarUmaRequisicaoGetPara(endpoint string) error {

	useCase := use_cases3.NovoListarClienteUseCase(ft.Repository)

	ft.Router.GET(endpoint, func(ctx *gin.Context) {
		controllers.ListarClientes(ctx, useCase)
	})

	req, err := http.NewRequest("GET", endpoint, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	ft.Response = httptest.NewRecorder()

	ft.Router.ServeHTTP(ft.Response, req)

	return nil
}

func (ft *FeatureTest) euEnviarUmaRequisioGetByIdPara(endpoint string) error {
	useCase := use_cases2.NovoBuscarClientePorIdUseCase(ft.Repository)

	ft.Router.GET(endpoint, func(ctx *gin.Context) {
		controllers.BuscarClientePorId(ctx, useCase, id)
	})

	req, err := http.NewRequest("GET", endpoint, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	ft.Response = httptest.NewRecorder()

	ft.Router.ServeHTTP(ft.Response, req)

	return nil
}

func (ft *FeatureTest) aRespostaDeveConterOsDetalhesDoCliente() error {
	var cliente use_cases4.ClienteOutput
	err := json.Unmarshal(ft.Response.Body.Bytes(), &cliente)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}
	return nil
}

func (ft *FeatureTest) queEuTenhoUmIdDeCliente() error {
	return nil
}

func (ft *FeatureTest) euEnviarUmaRequisioPUTPara(endpoint string) error {
	useCase := use_cases5.NovoAtualizarClienteUseCase(ft.Repository)

	ft.Router.PUT(endpoint, func(ctx *gin.Context) {
		controllers.AtualizarCliente(ctx, useCase, id)
	})

	body = strings.NewReader(`{"nome":"NovoPeão","dataNascimento":"2000-01-01","telefone":"123456789","endereco":"Test","email":"test@test.com"}`)

	req, err := http.NewRequest("PUT", endpoint, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	ft.Response = httptest.NewRecorder()

	ft.Router.ServeHTTP(ft.Response, req)

	return nil
}

func (ft *FeatureTest) queEuTenhoUmIdDeClienteEUmPayloadDeAtualizao() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext, t *testing.T) {
	ft := &FeatureTest{T: t}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		gin.SetMode(gin.TestMode)
		ft.Router = gin.Default()

		db := database.NewSqliteDB()
		ft.Repository = repositories.NovoClienteRepository(db)
		return ctx, nil
	})

	ctx.Step(`^que eu tenho um novo payload de cliente$`, ft.queEuTenhoUmNovoPayloadDeCliente)
	ctx.Step(`^eu enviar uma requisição POST para "([^"]*)"$`, ft.euEnviarUmaRequisicaoPostPara)
	ctx.Step(`^o código de resposta deve ser (\d+)$`, ft.oCodigoDeRespostaDeveSer)
	ctx.Step(`^a resposta deve conter um "([^"]*)"$`, ft.aRespostaDeveConterUm)
	ctx.Step(`^eu enviar uma requisição GET para "([^"]*)"$`, ft.euEnviarUmaRequisicaoGetPara)
	ctx.Step(`^o código de resposta deve ser (\d+)$`, ft.oCodigoDeRespostaDeveSer)
	ctx.Step(`^eu enviar uma requisição GetById para "([^"]*)"$`, ft.euEnviarUmaRequisioGetByIdPara)
	ctx.Step(`^a resposta deve conter os detalhes do cliente$`, ft.aRespostaDeveConterOsDetalhesDoCliente)
	ctx.Step(`^que eu tenho um id de cliente$`, ft.queEuTenhoUmIdDeCliente)
	ctx.Step(`^eu enviar uma requisição PUT para "([^"]*)"$`, ft.euEnviarUmaRequisioPUTPara)
	ctx.Step(`^que eu tenho um id de cliente e um payload de atualização$`, ft.queEuTenhoUmIdDeClienteEUmPayloadDeAtualizao)

}
