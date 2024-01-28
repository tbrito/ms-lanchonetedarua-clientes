# language: pt
Funcionalidade: API de Clientes
Como um usuário
Eu quero ser capaz de gerenciar clientes
Para que eu possa realizar operações CRUD

Cenário: Criar um novo cliente
Dado que eu tenho um novo payload de cliente
Quando eu enviar uma requisição POST para "/clientes"
Então o código de resposta deve ser 201
E a resposta deve conter um "id"

Cenário: Listar todos os clientes
Quando eu enviar uma requisição GET para "/clientes"
Então o código de resposta deve ser 200

Cenário: Obter um cliente por id
Dado que eu tenho um id de cliente
Quando eu enviar uma requisição GetById para "/clientes/{id}"
Então o código de resposta deve ser 200
E a resposta deve conter os detalhes do cliente

Cenário: Atualizar um cliente
Dado que eu tenho um id de cliente e um payload de atualização
Quando eu enviar uma requisição PUT para "/clientes/{id}"
Então o código de resposta deve ser 204