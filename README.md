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
go install [github.com/fullstorydev/grpcurl/cmd/grpcurl@latest](https://github.com/fullstorydev/grpcurl/cmd/grpcurl@latest)
