package ports

import "github.com/islanpedro01/microservices/shipping/internal/application/domain"

type APIPort interface {
	CalculateShipping(shipping domain.Shipping) (int32, error)
}