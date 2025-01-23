package orders

type OrderStatus string

const (
	OrderStatusWaiting   OrderStatus = "ожидает обработки логиста"
	OrderStatusActive    OrderStatus = "активен"
	OrderStatusCompleted OrderStatus = "завершен"
	OrderStatusPayed     OrderStatus = "оплачен"
)
