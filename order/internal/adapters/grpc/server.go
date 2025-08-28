package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/islanpedro01/microservices-proto/golang/order"
	"github.com/islanpedro01/microservices/order/config"
	"github.com/islanpedro01/microservices/order/internal/application/core/domain"
	"github.com/islanpedro01/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error){

	log.Printf("DEBUG: Requisição gRPC recebida: %+v", request)
    log.Printf("DEBUG: Número de itens na requisição: %d", len(request.OrderItems))

	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(int64(request.CostumerId), orderItems, request.TotalPrice)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{ OrderId: int32(result.ID)}, nil
}

type Adapter struct {
	api ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter (api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d, error: %v", a.port, err)
}
grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port")
	}
}