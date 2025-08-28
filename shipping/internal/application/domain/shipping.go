package domain

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	OrderID int64
	Items   []ShippingItem
}

func NewShipping(orderID int64, items []ShippingItem) Shipping {
	return Shipping{
		OrderID: orderID,
		Items:   items,
	}
}

func (s *Shipping) TotalQuantity() int32 {
	var total int32 = 0
	for _, item := range s.Items {
		total += item.Quantity
	}
	return total
}