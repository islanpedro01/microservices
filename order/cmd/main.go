package main

import (
	"log"
	"os"
	"github.com/islanpedro01/microservices/order/config"
	"github.com/islanpedro01/microservices/order/internal/adapters/db"
	"github.com/islanpedro01/microservices/order/internal/adapters/grpc"
	"github.com/islanpedro01/microservices/order/internal/adapters/payment_adapter"
	"github.com/islanpedro01/microservices/order/internal/adapters/shipping_adapter"
	"github.com/islanpedro01/microservices/order/internal/application/core/api"
	// "github.com/islanpedro01/microservices/order/internal/adapters/rest"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	shippingAdapter, err := shipping_adapter.NewAdapter(os.Getenv("SHIPPING_SERVICE_URL"))
	if err != nil {
		log.Fatalf("failed to initialize shipping adapter: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}