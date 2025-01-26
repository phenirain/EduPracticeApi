package orders

type OrderStatus string

const (
	OrderStatusReserved   OrderStatus = "зарезервирован"
	OrderStatusPaid       OrderStatus = "оплачен"
	OrderStatusDelivering OrderStatus = "в доставке"
	OrderStatusCompleted  OrderStatus = "завершен"
	OrderStatusCanceled   OrderStatus = "отменен"
)
