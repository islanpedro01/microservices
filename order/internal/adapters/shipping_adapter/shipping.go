package shipping_adapter

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/islanpedro01/microservices-proto/golang/shipping"
	"github.com/islanpedro01/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	}

	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) Create(ctx context.Context, order *domain.Order) (int32, error) {
	var shippingItems []*shipping.ShippingItem
	for _, item := range order.OrderItems {
		shippingItems = append(shippingItems, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	request := &shipping.CreateShippingRequest{
		OrderId: order.ID,
		Items:   shippingItems,
	}

	response, err := a.shipping.CreateShipping(ctx, request)
	if err != nil {
		return 0, err
	}

	return response.DeliveryDays, nil
}