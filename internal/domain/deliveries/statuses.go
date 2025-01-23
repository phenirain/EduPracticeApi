package deliveries

type DeliveryStatus string

const (
	DeliveryStatusOnTheWay  DeliveryStatus = "в пути"
	DeliveryStatusCanceled  DeliveryStatus = "отменен"
	DeliveryStatusCompleted DeliveryStatus = "завершен"
)
