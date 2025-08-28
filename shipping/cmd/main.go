package main

import (
	"log"
	"github.com/islanpedro01/microservices/shipping/internal/adapters/grpc"
	"github.com/islanpedro01/microservices/shipping/internal/application/core/api"

)

const (
	grpcPort = 9002
)

func main() {

	application := api.NewApplication()
	

	grpcAdapter := grpc.NewAdapter(application, grpcPort)

	log.Printf("Iniciando o serviço Shipping na porta %d", grpcPort)

	
	grpcAdapter.Run()
}

