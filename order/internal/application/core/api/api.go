package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/islanpedro01/microservices/order/internal/application/core/domain"
	"github.com/islanpedro01/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db: db,
		payment: payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	
		var productCodes []string
	for _, item := range order.OrderItems {
		productCodes = append(productCodes, item.ProductCode)
	}

	foundProducts, err := a.db.GetProductsByCodes(productCodes)
	if err != nil {
		return domain.Order{}, status.Errorf(codes.Internal, "erro ao consultar produtos: %v", err)
	}

	if len(foundProducts) != len(productCodes) {
		return domain.Order{}, status.Error(codes.NotFound, "um ou mais produtos não foram encontrados no estoque")
	}

	var totalItems int32 = 0
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}
	
	if totalItems > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "o número total de itens (%d) excede o limite de 50", totalItems)
	}
	
	err = a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	paymentErr := a.payment.Charge(ctx, &order)

	if paymentErr != nil {
		st, ok := status.FromError(paymentErr)
		if ok && st.Code() == codes.DeadlineExceeded {
			log.Printf("prazo para pagamento do pedido %d excedido. O serviço de pagamento demorou mais de 2 segundos para responder", order.ID)
		}
		order.Status = "Canceled"
		if updateErr := a.db.Update(&order); updateErr != nil {
		return domain.Order{}, fmt.Errorf("payment error: %v, update status error: %v", paymentErr, updateErr)
	}
	return domain.Order{}, paymentErr
}
order.Status = "Paid"
if updateErr := a.db.Update(&order); updateErr != nil {
	return domain.Order{}, updateErr
}
	shippingCtx, shippingCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shippingCancel()

	deliveryDays, shippingErr := a.shipping.Create(shippingCtx, &order)
	if shippingErr != nil {
		log.Printf("erro ao solicitar o cálculo do frete para o pedido %d: %v. Este erro não cancelará o pedido.", order.ID, shippingErr)
	} else {
		log.Printf("Pedido %d pago com sucesso. Prazo de entrega estimado em %d dias.", order.ID, deliveryDays)
	}

return order, nil
}