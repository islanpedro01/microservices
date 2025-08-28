package ports

import "github.com/islanpedro01/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	Update(*domain.Order) error
	GetProductsByCodes(codes []string) ([]domain.Product, error)
}

