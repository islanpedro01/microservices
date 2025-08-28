# Projeto de Microsserviços com gRPC e Go

Este repositório contém a implementação de um sistema de E-commerce simplificado baseado em uma arquitetura de microsserviços. O objetivo deste projeto é demonstrar os conceitos de comunicação síncrona usando gRPC, a aplicação da Arquitetura Hexagonal (Portas e Adaptadores) e a orquestração de múltiplos serviços com Docker Compose.

## Arquitetura e Tecnologias

O sistema é composto por três microsserviços principais:
* **Order**: Responsável por receber, validar e orquestrar o fluxo de criação de pedidos.
* **Payment**: Responsável por processar a cobrança de um pedido.
* **Shipping**: Responsável por calcular o prazo de entrega de um pedido.

**Tecnologias Utilizadas:**
* **Linguagem**: Go
* **Comunicação**: gRPC (com Protobuf)
* **Arquitetura**: Hexagonal (Portas e Adaptadores)
* **Banco de Dados**: MySQL
* **Containerização**: Docker e Docker Compose

## Pré-requisitos

Para executar este projeto, você precisará ter as seguintes ferramentas instaladas:
* [Docker](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/)
* [Go](https://golang.org/doc/install) (necessário para a ferramenta de teste)
* [grpcurl](https://github.com/fullstorydev/grpcurl) (ferramenta de linha de comando para interagir com serviços gRPC)

Você pode instalar o `grpcurl` com o seguinte comando Go:
```sh
go install [github.com/fullstorydev/grpcurl/cmd/grpcurl@latest](https://github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
Como Testar a Aplicação
Os testes são realizados através de chamadas gRPC diretas ao serviço Order usando o grpcurl.

Passo 1: Preparar o Banco de Dados
Antes de criar um pedido, é necessário cadastrar alguns produtos no estoque.

Abra um novo terminal e conecte-se ao container do banco de dados:

Bash

docker-compose exec db mysql -u root -p
(A senha é root)

Dentro do cliente MySQL, execute os seguintes comandos para criar a tabela e inserir os produtos:

SQL

USE microservices;

INSERT INTO products (created_at, updated_at, product_code, name, unit_price) VALUES
(NOW(), NOW(), 'P001', 'Teclado Gamer', 299.90),
(NOW(), NOW(), 'P002', 'Mouse sem fio', 120.50),
(NOW(), NOW(), 'P003', 'Monitor 24 polegadas', 899.99);

EXIT;
Passo 2: Executar Testes com grpcurl
Execute os comandos abaixo no seu terminal para simular diferentes cenários.

Teste 1: Cenário de Sucesso (Pedido Válido)
Bash

grpcurl -plaintext -d '{ "costumer_id": 123, "order_items": [ { "product_code": "P001", "quantity": 2, "unit_price": 299.90 }, { "product_code": "P002", "quantity": 1, "unit_price": 120.50 } ], "total_price": 720.30 }' localhost:9001 Order/Create
Resultado Esperado: Um JSON com o ID do pedido. Ex: { "order_id": 1 }. Nos logs, você verá o fluxo completo de comunicação entre os serviços.

Teste 2: Erro - Produto Não Encontrado
Bash

grpcurl -plaintext -d '{ "costumer_id": 124, "order_items": [ { "product_code": "P999", "quantity": 1, "unit_price": 10.00 } ], "total_price": 10.00 }' localhost:9001 Order/Create
Resultado Esperado: Um erro com code = NotFound e a mensagem "um ou mais produtos não foram encontrados no estoque".

Teste 3: Erro - Pagamento Recusado (Valor > R$ 1000)
Bash

grpcurl -plaintext -d '{ "costumer_id": 125, "order_items": [ { "product_code": "P003", "quantity": 2, "unit_price": 899.99 } ], "total_price": 1799.98 }' localhost:9001 Order/Create
Resultado Esperado: Um erro com code = InvalidArgument e a mensagem de que pagamentos acima de 1000 não são permitidos.

Teste 4: Erro - Limite de Itens Excedido (> 50)
Bash

grpcurl -plaintext -d '{ "costumer_id": 126, "order_items": [ { "product_code": "P001", "quantity": 51, "unit_price": 299.90 } ], "total_price": 15294.9 }' localhost:9001 Order/Create
Resultado Esperado: Um erro com code = InvalidArgument e a mensagem de que o limite de 50 itens foi excedido.
