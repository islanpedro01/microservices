package api

import (
	"github.com/islanpedro01/microservices/shipping/internal/application/domain"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) CalculateShipping(shipping domain.Shipping) (int32, error) {
	totalQuantity := shipping.TotalQuantity()

	deliveryDays := 1 + (totalQuantity / 5)

	return deliveryDays, nil
}