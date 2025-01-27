package orders

type OrderStatus string

const (
	OrderStatusReserved   OrderStatus = "Reserved"
	OrderStatusPaid       OrderStatus = "Paid"
	OrderStatusDelivering OrderStatus = "Delivering"
	OrderStatusCompleted  OrderStatus = "Completed"
	OrderStatusCanceled   OrderStatus = "Canceled"
)
