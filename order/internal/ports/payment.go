package ports

import (
	"context"

	"github.com/islanpedro01/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
}