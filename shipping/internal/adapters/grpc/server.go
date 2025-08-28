package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/islanpedro01/microservices-proto/golang/shipping"
	"github.com/islanpedro01/microservices/shipping/internal/application/domain"
	"github.com/islanpedro01/microservices/shipping/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, a)
	
	// Para ambiente de desenvolvimento, para permitir a exploração do serviço.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a *Adapter) CreateShipping(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	var shippingItems []domain.ShippingItem
	for _, item := range request.Items {
		shippingItems = append(shippingItems, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingDomain := domain.NewShipping(request.OrderId, shippingItems)

	deliveryDays, err := a.api.CalculateShipping(shippingDomain)
	if err != nil {
		return nil, err
	}

	return &shipping.CreateShippingResponse{DeliveryDays: deliveryDays}, nil
}