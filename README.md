# MS de Gerenciamento de Clientes

Esta é uma API baseada em Go para gerenciar clientes. Ela fornece endpoints para criar, ler, atualizar e deletar dados de clientes.

## Começando

Estas instruções permitirão que você obtenha uma cópia do projeto e a execute em sua máquina local para fins de desenvolvimento e teste.

### Pré-requisitos

- Go (última versão)
- Docker (última versão)

### Instalando

1. Clone o repositório para sua máquina local.
2. Navegue até o diretório do projeto.
3. Execute `go mod download` para baixar os módulos Go necessários.
4. Construa a imagem Docker executando `docker build -t my-go-app .`
5. Execute o contêiner Docker com `docker run -p 8080:8080 my-go-app`

## Endpoints da API

- `GET /clientes`: Buscar todos os clientes.
- `GET /clientes/{id}`: Buscar um único cliente por ID.
- `POST /clientes`: Criar um novo cliente.
- `PUT /clientes/{id}`: Atualizar um cliente existente.


## Construído com

- [Go](https://golang.org/) - A linguagem de programação utilizada.
- [Docker](https://www.docker.com/) - Usado para a contêinerização.

## Autores

- **janduymonroe** - *Trabalho inicial* - [janduymonroe](https://github.com/janduymonroe)
- **tbrito** - *Trabalho inicial* - [tbrito](https://github.com/tbrito)

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE.md](LICENSE.md) para detalhes